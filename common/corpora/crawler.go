package corpora

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

const websiteURL = "https://wortschatz.uni-leipzig.de/en/download"

func listLanguages() (res []string, err error) {
	res = make([]string, 0)
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)
	c.OnHTML(`a.btn-modal[data-container=body][data-toggle=tooltip][href]`, func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		res = append(res, link)
	})
	err = c.Visit(websiteURL)
	return
}

func (c *Corpora) loadFromWebsite(onStart func(int), onLoad func()) error {
	languages, err := listLanguages()
	if err != nil {
		return err
	}
	onStart(len(languages))
	collector := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
		colly.Async(true),
	)
	if err := collector.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: c.Parallelism}); err != nil {
		return nil
	}
	collector.OnHTML("html", func(e *colly.HTMLElement) {
		language := strings.ReplaceAll(e.ChildText("title"), "Download Corpora ", "")

		languageCorpuses := &Corpuses{Corpuses: make([]Corpuse, 0)}

		e.ForEachWithBreak("table.table", func(_ int, e *colly.HTMLElement) bool {
			e.ForEach("tr[id^=corpora_type_]", func(_ int, c *colly.HTMLElement) {
				typeID := strings.ReplaceAll(c.Attr("id"), "corpora_type_", "")
				typeName := strings.TrimSpace(c.ChildText("th"))
				listCorpuses(typeName, typeID, e, &languageCorpuses.Corpuses)
			})
			return true
		})

		c.index.mu.Lock()
		c.index.Corpuses[language] = languageCorpuses
		c.index.mu.Unlock()
		// p.Increment()
		onLoad()
	})
	for _, v := range languages {
		_ = collector.Visit(v)
	}
	collector.Wait()
	return nil
}

func listCorpuses(name, typeID string, e *colly.HTMLElement, corpuses *[]Corpuse) {
	e.ForEach(fmt.Sprintf("tr[id*=_%s_]", typeID), func(_ int, c *colly.HTMLElement) {
		country := ""
		if c.DOM.Children().Length() == 3 {
			country = strings.TrimSpace(c.ChildText("td:nth-child(2)"))
		}
		year := c.ChildText("td:first-child")
		options := make([]Option, 0)
		c.ForEach("a.link_corpora_download", func(_ int, e *colly.HTMLElement) {
			id := e.Attr("id")
			size := e.Text
			options = append(options, Option{ID: id, Size: size})
		})
		*corpuses = append(*corpuses, Corpuse{Name: name, Year: year, Country: country, Options: options})
	})
}

func (c *Corpuses) Save(name string) (err error) {
	var f *os.File
	var d []byte

	if f, err = os.Create(name); err != nil {
		return
	}
	defer f.Close()

	if d, err = json.Marshal(c); err != nil {
		return
	}
	_, err = f.Write(d)
	return
}
