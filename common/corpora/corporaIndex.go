package corpora

import (
	"fmt"
	"io/fs"
	"os"
	"slices"
	"sync"
	"time"
)

const day = time.Duration(24) * time.Hour

type CorporaIndex struct {
	Corpuses map[string]*Corpuses `yaml:",inline"`
	mu       sync.Mutex
}

type Corpuses struct {
	Corpuses []Corpuse
}

type Corpuse struct {
	Name    string
	Country string
	Year    string
	Options []Option
}

type Option struct {
	ID   string
	Size string
}

type display struct {
	availabledOptions map[string]string
	options           []string
}

func (d *display) Options() []string {
	return d.options
}

func (d *display) AvailabledOptions() map[string]string {
	return d.availabledOptions
}

func (c *Corpora) Languages() []string {
	return c.index.languages()
}

func (c *Corpora) Corpuses(name string) *Corpuses {
	corpuses, ok := c.index.Corpuses[name]
	if !ok {
		return nil
	}
	return corpuses
}

func (c *Corpora) Load(onNew func(), onStart func(int), onLoad func(), onEnd func()) error {
	index, err := tryLoadFromCache(c.CacheFile, c.Force)
	if err != nil {
		return err
	}
	c.index = index
	isNew := index == nil
	if isNew {
		c.index = &CorporaIndex{Corpuses: make(map[string]*Corpuses)}
		onNew()
		err := c.loadFromWebsite(onStart, onLoad)
		onEnd()
		if err != nil {
			return err
		}
	}
	return c.index.marshal(c.CacheFile)
}

func tryLoadFromCache(cacheFile string, force bool) (index *CorporaIndex, err error) {
	if force {
		return
	}
	var stat fs.FileInfo
	stat, err = os.Stat(cacheFile)
	if err != nil {
		if os.IsNotExist(err) {
			err = nil
		}
		return
	}
	if stat.ModTime().Add(day).Before(time.Now()) {
		return
	}

	return unMarshal(cacheFile)
}

func (c *CorporaIndex) languages() []string {
	res := make([]string, len(c.Corpuses))
	i := 0
	for k := range c.Corpuses {
		res[i] = k
		i++
	}
	slices.Sort(res)
	return res
}

func (c *Corpuses) Display() (res display) {
	res = display{options: make([]string, 0), availabledOptions: make(map[string]string)}
	for _, c := range c.Corpuses {
		country := c.Country
		if country != "" {
			country += " "
		}
		baseName := fmt.Sprintf("%s: %s %s", c.Name, c.Year, country)
		for _, o := range c.Options {
			name := fmt.Sprintf("%s %s", baseName, o.Size)
			res.options = append(res.options, name)
			res.availabledOptions[name] = o.ID
		}
	}

	// slices.Sort(res.options)
	return
}
