package config

import (
	"fmt"
)

func New() *Config {
	return &Config{data: &ConfigData{}}
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
	Languages map[string]LayoutLanguage
}

type LayoutLanguage struct {
	Weight float64
	Scores string
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
