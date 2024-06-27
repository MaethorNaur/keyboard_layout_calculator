package corpora

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/pterm/pterm"
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

func (c *Corpora) loadFromWebsite() error {
	multi := &pterm.DefaultMultiPrinter
	_, _ = multi.Start()
	defer func() { _, _ = multi.Stop() }()
	spinner, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start("Listing from Corpora")
	languages, err := listLanguages()
	if err != nil {
		return err
	}
	p, _ := pterm.DefaultProgressbar.WithWriter(multi.NewWriter()).WithTotal(len(languages)).WithTitle("Languages retrieved").Start()
	spinner.UpdateText("Retrieving corpuses from Corpora")
	collector := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
		colly.Async(true),
	)
	if err := collector.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: c.Parallelism}); err != nil {
		return nil
	}
	collector.OnHTML("html", func(e *colly.HTMLElement) {
		language := strings.ReplaceAll(e.ChildText("title"), "Download Corpora ", "")

		languageCorpuses := &Corpuses{Corpuses: make(map[string][]Corpuse)}

		e.ForEachWithBreak("table.table", func(_ int, e *colly.HTMLElement) bool {
			e.ForEach("tr[id^=corpora_type_]", func(_ int, c *colly.HTMLElement) {
				typeID := strings.ReplaceAll(c.Attr("id"), "corpora_type_", "")
				typeName := strings.TrimSpace(c.ChildText("th"))
				corpuses := listCorpuses(typeID, e)
				languageCorpuses.Corpuses[typeName] = corpuses
			})
			return true
		})

		c.index.mu.Lock()
		c.index.Corpuses[language] = languageCorpuses
		c.index.mu.Unlock()
		p.Increment()
	})
	for _, v := range languages {
		_ = collector.Visit(v)
	}
	collector.Wait()
	spinner.Info("Corpuses retrieved")
	return nil
}

func listCorpuses(typeID string, e *colly.HTMLElement) []Corpuse {
	res := make([]Corpuse, 0)
	e.ForEach(fmt.Sprintf("tr[id*=_%s_]", typeID), func(_ int, c *colly.HTMLElement) {
		country := ""
		if c.DOM.Children().Length() == 3 {
			country = strings.TrimSpace(c.ChildText("td:nth-child(2)"))
		}
		year, _ := strconv.Atoi(c.ChildText("td:first-child"))
		options := make([]Option, 0)
		c.ForEach("a.link_corpora_download", func(_ int, e *colly.HTMLElement) {
			id := e.Attr("id")
			size := e.Text
			options = append(options, Option{ID: id, Size: size})
		})
		res = append(res, Corpuse{Year: year, Country: country, Options: options})
	})
	return res
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
