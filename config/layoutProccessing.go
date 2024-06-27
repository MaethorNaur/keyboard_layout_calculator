package config

import (
	"slices"

	"github.com/StudioSol/set"
)

type LayoutProcessing struct {
	layouts map[string][]languages
}

type languages struct {
	language *ConfigLanguage
	options  LayoutLanguage
}

func (config *Config) newLayoutProcessing() (l *LayoutProcessing) {
	l = &LayoutProcessing{layouts: make(map[string][]languages)}
	for _, layout := range config.data.Layouts {
		ll := make([]languages, 0)
		for name, l := range layout.Languages {
			if lang, err := config.Language(name); err == nil {
				ll = append(ll, languages{language: lang, options: l})
			}
		}
		l.layouts[layout.Name] = ll
	}
	return
}

func (l *LayoutProcessing) filter(layouts []string) {
	if len(layouts) == 0 {
		return
	}

	layoutsSet := set.NewLinkedHashSetString(layouts...)
	for k := range l.layouts {
		if !layoutsSet.InArray(k) {
			delete(l.layouts, k)
		}
	}
}

func (l *LayoutProcessing) Names() (r []string) {
	r = make([]string, len(l.layouts))
	i := 0
	for k := range l.layouts {
		r[i] = k
		i++
	}
	slices.Sort(r)
	return
}

func (l *LayoutProcessing) Layouts() map[string][]languages {
	return l.layouts
}

func (l *languages) Language() *ConfigLanguage {
	return l.language
}

func (l *languages) Options() LayoutLanguage {
	return l.options
}
