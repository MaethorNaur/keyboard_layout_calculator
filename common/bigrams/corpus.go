package bigrams

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"path/filepath"
	"strings"
)

const urlPrefix = "https://downloads.wortschatz-leipzig.de/corpora/"

func humanizeBytes(s float64) (value, size string) {
	sizes := []string{" B", " kB", " MB", " GB", " TB", " PB", " EB"}
	base := 1024.0
	if s < 10 {
		return fmt.Sprintf("%2.0f", s), sizes[0]
	}
	e := math.Floor(logn(s, base))
	suffix := sizes[int(e)]
	val := math.Floor(s/math.Pow(base, e)*10+0.5) / 10
	f := "%.0f"
	if val < 10 {
		f = "%.1f"
	}

	return fmt.Sprintf(f, val), suffix
}

func logn(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}

func download(task task, name string) (data []byte, err error) {
	var uncompressedStream *gzip.Reader
	var req *http.Request
	var resp *http.Response

	if req, err = http.NewRequestWithContext(context.Background(),
		http.MethodGet,
		fmt.Sprintf("%s%s.tar.gz", urlPrefix, name),
		http.NoBody); err != nil {
		return
	}
	if resp, err = http.DefaultClient.Do(req); err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", resp.Status)
		return
	}
	b, s := humanizeBytes(float64(resp.ContentLength))

	task.callbacks.onFetch(Downloading,
		task.language.Name(),
		fmt.Sprintf("%s.tar.gz", name),
		fmt.Sprintf("%s %s", b, s))

	if uncompressedStream, err = gzip.NewReader(resp.Body); err != nil {
		return
	}

	tarReader := tar.NewReader(uncompressedStream)

	for {
		var header *tar.Header
		if header, err = tarReader.Next(); err != nil {
			if err == io.EOF {
				err = fmt.Errorf("word file not found for: %s", name)
			}
			return
		}

		if header.Typeflag == tar.TypeReg {
			if _, filename := filepath.Split(header.Name); strings.HasSuffix(filename, "words.txt") {
				buf := new(bytes.Buffer)
				b, s := humanizeBytes(float64(header.Size))
				task.callbacks.onFetch(Extracting, task.language.Name(), filename, fmt.Sprintf("%s %s", b, s))

				if _, err = buf.ReadFrom(tarReader); err == nil {
					data = buf.Bytes()
				}
				return
			}
		}
	}
}
