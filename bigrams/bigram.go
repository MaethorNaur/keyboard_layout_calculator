package bigrams

import (
	"bufio"
	"bytes"
	"cmp"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode/utf8"

	"slices"
)

type NGram struct {
	Md5                string
	Monograms          []NGramEntry
	Bigrams            []NGramEntry
	MonogramsInBigrams []NGramEntry
	// MonogramsInBigrams map[string]int
}

type NGramEntry struct {
	Value     string
	Frequency float64
}

func NewNGram() *NGram {
	return &NGram{
		Monograms:          make([]NGramEntry, 0),
		Bigrams:            make([]NGramEntry, 0),
		MonogramsInBigrams: make([]NGramEntry, 0),
		// MonogramsInBigrams: make(map[string]int),
	}
}

func (r *Runner) load(task task) (n *NGram, err error) {
	n = NewNGram()
	hasher := md5.New()
	for _, corpuse := range task.language.Corpuses {
		hasher.Write([]byte(corpuse))
	}
	hasher.Write([]byte(task.language.Letters))

	file := filepath.Join(r.cacheDir, fmt.Sprintf("%s.json", task.language.Name()))

	_, errState := os.Stat(file)
	downloadCorpuses := os.IsNotExist(errState) || r.force
	md5Value := hex.EncodeToString(hasher.Sum(nil))

	if !os.IsNotExist(errState) {
		if err = n.read(file); err != nil {
			return
		}
		if md5Value != n.Md5 {
			downloadCorpuses = true
		}
	}

	if downloadCorpuses {
		n.Md5 = md5Value
		data := make([]byte, 0)
		for _, name := range task.language.Corpuses {
			var d []byte
			if d, err = download(task, name); err != nil {
				continue
			}
			data = append(data, d...)
			data = append(data, "\n"...)
		}
		n.compute(data, task)
		err = n.save(file)
	}
	return
}

func (n *NGram) read(file string) (err error) {
	var f *os.File
	var d []byte
	if f, err = os.Open(file); err != nil {
		return
	}
	defer f.Close()
	if d, err = io.ReadAll(f); err != nil {
		return
	}
	err = json.Unmarshal(d, &n)
	return
}

func reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func (n *NGram) compute(d []byte, task task) {
	task.spinner.UpdateText(fmt.Sprintf("[%s] Computing bigrams", task.language.Name()))
	re := regexp.MustCompile(fmt.Sprintf("^[%s]+", task.language.Letters))

	r := bufio.NewReader(bytes.NewReader(d))

	monograms := make(map[string]float64)
	bigrams := make(map[string]float64)
	monogramsInBigrams := make(map[string]float64)
	for {
		w := 0
		data, _, err := r.ReadLine()
		if err != nil {
			break
		}
		line := string(data)
		split := strings.Split(line, "\t")
		if len(split) < 3 {
			continue
		}
		word := strings.ToLower(split[1])
		wordLen := len(word)
	out:
		for i := 0; i < wordLen; i += w {
			var buf []byte
			pos := i
			hasBigram := true
			for j := 0; j < 2; j++ {
				if pos >= wordLen {
					break out
				}
				runeValue, width := utf8.DecodeRuneInString(word[pos:])
				pos += width
				w += width
				value := string(runeValue)

				if !re.MatchString(value) {
					hasBigram = false
					continue
				}
				if _, ok := monograms[value]; !ok {
					monograms[value] = 0
					monogramsInBigrams[value] = 0
				}
				monograms[value] += 1
				buf = utf8.AppendRune(buf, runeValue)
			}
			if hasBigram {
				var value = string(buf)
				size := len(value)
				for start := 0; start < size; {
					r, rn := utf8.DecodeRuneInString(value[start:])
					start += rn
					s := string(r)
					monogramsInBigrams[s] += 1
				}
				reversedValue := reverse(value)
				if _, ok := bigrams[reversedValue]; ok {
					value = reversedValue
				}

				if _, ok := bigrams[value]; !ok {
					bigrams[value] = 0
				}
				bigrams[value] += 1
			}
		}
	}
	frequencySum := 0.
	for _, v := range monograms {
		frequencySum += v
	}
	for k := range monograms {
		monograms[k] /= frequencySum
	}
	size := len(task.language.Letters)
	for start := 0; start < size; {
		r, rn := utf8.DecodeRuneInString(task.language.Letters[start:])
		start += rn
		s := string(r)
		if _, ok := monogramsInBigrams[s]; !ok {
			monogramsInBigrams[s] = 0
		}
		if _, ok := monograms[s]; !ok {
			monograms[s] = 0
		}
	}
	n.Monograms = mapToEntry(monograms)
	n.MonogramsInBigrams = mapToEntry(monogramsInBigrams)
	n.Bigrams = mapToEntry(bigrams)
	n.sort()
}

func (n *NGram) sort() {
	slices.SortFunc(n.Monograms, sort)
	slices.SortFunc(n.MonogramsInBigrams, sort)
	slices.SortFunc(n.Bigrams, sort)
}

func sort(a, b NGramEntry) int {
	res := cmp.Compare(b.Frequency, a.Frequency)
	if res == 0 {
		return cmp.Compare(b.Value, a.Value)
	}
	return res
}

func (n *NGram) Merge(o *NGram, weight float64) {
	if weight == 0 {
		return
	}
	monograms := entryToMap(n.Monograms)
	bigrams := entryToMap(n.Bigrams)
	monogramsInBigrams := entryToMap(n.MonogramsInBigrams)
	merge(monograms, o.Monograms, weight)
	merge(bigrams, o.Bigrams, weight)
	merge(monogramsInBigrams, o.MonogramsInBigrams, weight)
	n.Monograms = mapToEntry(monograms)
	n.Bigrams = mapToEntry(bigrams)
	n.MonogramsInBigrams = mapToEntry(monogramsInBigrams)
	n.sort()
}

func entryToMap(entries []NGramEntry) (res map[string]float64) {
	res = make(map[string]float64)

	for _, entry := range entries {
		res[entry.Value] = entry.Frequency
	}
	return
}

func mapToEntry(m map[string]float64) (res []NGramEntry) {
	res = make([]NGramEntry, 0)
	for k, v := range m {
		res = append(res, NGramEntry{Value: k, Frequency: v})
	}
	return
}

func merge(m map[string]float64, entries []NGramEntry, weight float64) {
	percent := weight / 100.

	for _, entry := range entries {
		if _, ok := m[entry.Value]; !ok {
			m[entry.Value] = 0
		}
		m[entry.Value] += entry.Frequency * percent
	}
}

func (n *NGram) save(name string) (err error) {
	var f *os.File
	var d []byte

	if f, err = os.Create(name); err != nil {
		return
	}
	defer f.Close()

	if d, err = json.Marshal(n); err != nil {
		return
	}
	_, err = f.Write(d)
	return
}
