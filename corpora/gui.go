package corpora

import (
	"fmt"

	"atomicgo.dev/keyboard/keys"
	"github.com/pterm/pterm"
)

func (c *Corpora) SelectCorpuses() (corpuseIDs []string) {
	area, _ := pterm.DefaultArea.WithFullscreen().Start()
	defer func() {
		_ = area.Stop()
	}()
	corpuseIDs = make([]string, 0)
	language, _ := pterm.DefaultInteractiveSelect.
		WithOptions(c.index.languages()).
		// WithMaxHeight(pterm.GetTerminalHeight()).
		WithDefaultText("Select a language").
		Show()
	corpuses, ok := c.index.Corpuses[language]
	if !ok {
		return
	}
	display := corpuses.display()
	corpuseIDs, _ = pterm.DefaultInteractiveMultiselect.
		WithOptions(display.options).
		WithDefaultText(fmt.Sprintf("Select one or more corpuses for %s", language)).
		WithKeyConfirm(keys.Enter).
		WithKeySelect(keys.Space).
		WithMaxHeight(pterm.GetTerminalHeight()).
		Show()
	for i, v := range corpuseIDs {
		corpuseIDs[i] = display.availabledOptions[v]
	}
	return
}
