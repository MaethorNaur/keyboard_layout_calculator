package config

import (
	"github.com/pterm/pterm"
)

func (config *Config) SelectLayouts(layouts []string) (l *LayoutProcessing) {
	l = config.newLayoutProcessing()
	if len(layouts) == 0 {
		layouts, _ = pterm.DefaultInteractiveMultiselect.
			WithOptions(l.Names()).
			// WithDefaultText("Select which layout(s) to generate").
			// WithKeyConfirm(keys.Enter).
			// WithKeySelect(keys.Space).
			Show()
	} else {
		for _, v := range layouts {
			if v == "*" {
				layouts = make([]string, 0)
				break
			}
		}
	}
	l.filter(layouts)
	return
}

func (config *Config) SelectLanguage() (l *ConfigLanguage, err error) {
	options := make([]string, len(config.data.Languages))
	i := 0
	for name := range config.data.Languages {
		options[i] = name
		i++
	}
	var selected string
	selected, err = pterm.DefaultInteractiveSelect.WithOptions(options).Show()
	if err == nil {
		l = config.data.Languages[selected]
		l.name = selected
	}
	return

}
