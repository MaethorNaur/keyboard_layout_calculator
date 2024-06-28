package layout

import (
	"fmt"
	"strings"

	"github.com/8VIM/keyboard_layout_calculator/common/config"
	"github.com/pterm/pterm"
)

type LayoutCmd struct {
	Add AddCmd `cmd:"" help:"Add a new layout"`
}

type AddCmd struct {
	Name      string   `short:"n" optional:"" help:"Layout name"`
	Languages []string `arg:"" optional:"" name:"language"`
}

func (c AddCmd) Run(conf *config.Config) error {
	name := strings.TrimSpace(c.Name)
	if name == "" {
		name, _ = pterm.DefaultInteractiveTextInput.Show("Layout name")
		name = strings.TrimSpace(name)
		if name == "" {
			return fmt.Errorf("layout name should not be empty")
		}
	}

	name = strings.ToLower(name)
	layout := conf.Layout(name)

	if layout != nil {
		return fmt.Errorf("Layout: %s already exists", name)
	}

	layout = conf.NewLayout(name)

	if len(c.Languages) == 0 {
		languages := conf.Languages()
		if len(languages) == 0 {
			return fmt.Errorf("At least a language must be defined")
		}

		selected, _ := pterm.DefaultInteractiveMultiselect.
			WithOptions(languages).
			WithMaxHeight(pterm.GetTerminalHeight()).
			Show()

		c.Languages = selected
	}

	options := make([]string, 0)

	for _, l := range c.Languages {
		l = strings.ToLower(l)
		if conf.HasLanguage(l) {
			options = append(options, l)
		}
	}

	if len(options) == 0 {
		return fmt.Errorf("At least a language must be defined")
	}
	for _, l := range options {
		layout.AddLanguage(l)
	}
	return conf.Save()
}
