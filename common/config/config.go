package config

import (
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/8VIM/keyboard_layout_calculator/common/internal/utils"
)

func New() *Config {
	return &Config{data: &ConfigData{Languages: map[string]*ConfigLanguage{},
		Layouts: make([]*Layout, 0),
		Scores:  map[string][][]float64{},
	}}
}

type Config struct {
	File string
	data *ConfigData
}

type ConfigData struct {
	Languages map[string]*ConfigLanguage
	Layouts   []*Layout
	Scores    map[string][][]float64
}

type Layout struct {
	Name      string
	Languages map[string]*LayoutLanguage
}

type LayoutLanguage struct {
	Weight float64
	Scores string
}

func (c *Config) Init(path string) error {
	c.File = filepath.Clean(path)
	dir, file := filepath.Split(path)
	info, err := os.Stat(c.File)
	if file == "" || (err == nil && info.IsDir()) {
		return fmt.Errorf("%s should not be a directory", c.File)
	}
	if errors.Is(err, fs.ErrNotExist) {
		slog.Info(fmt.Sprintf("%s doesn't exists. Preparing directories for it", c.File))
		if err := utils.EnsureDirExists(filepath.Clean(dir)); err != nil {
			return err
		}
	}
	c.Load()
	return nil
}

func (config *Config) Language(name string) (l *ConfigLanguage, err error) {
	var ok bool

	if l, ok = config.data.Languages[name]; !ok {
		err = fmt.Errorf("'%s' is not defined in the current configuration", name)
	} else {
		l.name = name
	}
	return
}

func (config *Config) BatchEdit() *batchEdit {
	return &batchEdit{actions: make([]editAction, 0), config: config}
}

func (config *Config) HasLanguage(name string) (ok bool) {
	_, ok = config.data.Languages[name]
	return
}

func (config *Config) Languages() []string {
	languages := make([]string, 0)
	for name := range config.data.Languages {
		languages = append(languages, name)
	}
	return languages
}

func (config *Config) NewLayout(name string) *Layout {
	layout := &Layout{Name: name, Languages: map[string]*LayoutLanguage{}}
	config.data.Layouts = append(config.data.Layouts, layout)
	return layout
}

func (config *Config) Layout(name string) *Layout {
	for _, l := range config.data.Layouts {
		if l.Name == name {
			return l
		}
	}
	return nil
}

func (l *Layout) AddLanguage(name string) {
	for ln := range l.Languages {
		if ln == name {
			return
		}
	}

	l.Languages[name] = &LayoutLanguage{Weight: 100, Scores: "kjoetom"}
}
