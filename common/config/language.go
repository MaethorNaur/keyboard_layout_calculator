package config

import (
	"fmt"
	"log/slog"
)

type ConfigLanguage struct {
	Corpuses []string
	Letters  string
	name     string
}

func (c *ConfigLanguage) SetName(name string) {
	c.name = name
}

func (config *Config) LanguagesSize() int {
	return len(config.data.Languages)
}

func (config *Config) AddLanguage(name, letters string) *Config {
	if config.data.Languages == nil {
		config.data.Languages = map[string]*ConfigLanguage{}
	}
	config.data.Languages[name] = newLanguage(name, letters)
	return config
}

func newLanguage(name, letters string) *ConfigLanguage {
	return &ConfigLanguage{Corpuses: make([]string, 0), Letters: letters, name: name}
}

func (l *ConfigLanguage) Name() string {
	return l.name
}

func (l *ConfigLanguage) RemoveCorpuse(name string) bool {
	msg := "'%s' is not part of the corpuses of '%s'"
	defer func() { slog.Debug(fmt.Sprintf(msg, name, l.name)) }()
	for i, e := range l.Corpuses {
		if e == name {
			l.Corpuses = append(l.Corpuses[:i], l.Corpuses[i+1:]...)
			msg = "'%s' have been removed to the corpuses of '%s'"
			return true
		}
	}
	return false
}

func (l *ConfigLanguage) AddCorpuse(name string) bool {
	msg := "'%s' have been added to the corpuses of '%s'"
	defer func() { slog.Debug(fmt.Sprintf(msg, name, l.name)) }()
	for _, e := range l.Corpuses {
		if e == name {
			msg = "'%s' is already part of the corpuses of '%s'"
			return false
		}
	}
	l.Corpuses = append(l.Corpuses, name)
	return true
}
