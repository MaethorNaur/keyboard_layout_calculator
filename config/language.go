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

func (l *ConfigLanguage) Name() string {
	return l.name
}

func (l *ConfigLanguage) RemoveCorpuse(name string) bool {
	msg := "'%s' is not part of the corpuses of '%s'"
	defer func() { slog.Debug(fmt.Sprintf(msg, name, l.name)) }()
	for i, e := range l.Corpuses {
		if e == name {
			l.Corpuses = append(l.Corpuses[:i], l.Corpuses[:i+1]...)
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
