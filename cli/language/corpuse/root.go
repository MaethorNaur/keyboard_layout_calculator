package corpuse

import (
	"fmt"
	"log/slog"

	"github.com/8VIM/keyboard_layout_calculator/common/config"
	"github.com/8VIM/keyboard_layout_calculator/common/corpora"
	"github.com/pterm/pterm"
)

type CorpuseCmd struct {
	Add    AddCmd    `cmd:"" help:"Add new corpuse(s) if there are not already defined. If 'corpuse' is not provided you will be able to select corpuses from 'Corpora Collection Leipzig'"`
	Remove RemoveCmd `cmd:"" help:"Remove the given corpuses. All non matching 'corpuse' for the language will be ignored. If 'corpuse' is not provided you will be able to select from the currently ones defined."`
}

type AddCmd struct {
	Corpuses []string `arg:"" optional:"" name:"corpuse"`
}
type RemoveCmd struct {
	Corpuses []string `arg:"" optional:"" name:"corpuse"`
}

func (c AddCmd) Run(cache *corpora.Corpora, language *config.ConfigLanguage, conf *config.Config) error {
	if len(c.Corpuses) == 0 {
		c.Corpuses = selectCorpuses(cache)
	}
	be := conf.BatchEdit()
	for _, e := range c.Corpuses {
		be.Add(func(_ *config.Config) bool {
			return language.AddCorpuse(e)
		})
	}
	if _, err := be.Execute(); err != nil {
		slog.Error(err.Error())
	}
	return nil
}

func (c RemoveCmd) Run(language *config.ConfigLanguage, conf *config.Config) error {
	if len(c.Corpuses) == 0 {
		c.Corpuses, _ = pterm.DefaultInteractiveMultiselect.
			WithOptions(language.Corpuses).
			Show()
	}
	be := conf.BatchEdit()
	for _, e := range c.Corpuses {
		be.Add(func(_ *config.Config) bool {
			return language.RemoveCorpuse(e)
		})
	}
	if _, err := be.Execute(); err != nil {
		slog.Error(err.Error())
	}
	return nil
}

func selectCorpuses(c *corpora.Corpora) (corpuseIDs []string) {
	area, _ := pterm.DefaultArea.WithFullscreen().Start()
	defer func() {
		_ = area.Stop()
	}()
	corpuseIDs = make([]string, 0)
	language, _ := pterm.DefaultInteractiveSelect.
		WithOptions(c.Languages()).
		WithMaxHeight(pterm.GetTerminalHeight()).
		WithDefaultText("Select a language for the corpuses").
		Show()
	corpuses := c.Corpuses(language)
	if corpuses == nil {
		return
	}
	display := corpuses.Display()
	corpuseIDs, _ = pterm.DefaultInteractiveMultiselect.
		WithOptions(display.Options()).
		WithDefaultText(fmt.Sprintf("Select one or more corpuses for %s", language)).
		// WithKeyConfirm(keys.Enter).
		// WithKeySelect(keys.Space).
		WithMaxHeight(pterm.GetTerminalHeight()).
		Show()
	for i, v := range corpuseIDs {
		corpuseIDs[i] = display.AvailabledOptions()[v]
	}
	return
}
