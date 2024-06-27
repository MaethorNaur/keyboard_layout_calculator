package corpuse

import (
	"log/slog"

	"github.com/8VIM/keyboard_layout_calculator/config"
	"github.com/8VIM/keyboard_layout_calculator/corpora"
	"github.com/pterm/pterm"
)

type CorpuseCmd struct {
	Add    AddCmd    `cmd:"" help:"Add new corpuse(s) if there are not already defined. If 'corpuse' is not provided you will be able to select corpuses from 'Corpora Collection Leipzig'"`
	Remove RemoveCmd `cmd:"" help:"Remove the given corpuses. All non matching 'corpuse' for the language will be ignored. If 'corpuse' is not provided you will be able to select from the currently ones defined."`
}

type AddCmd struct {
	Corpuse []string `arg:"" optional:""`
}
type RemoveCmd struct {
	Corpuse []string `arg:"" optional:""`
}

func (c AddCmd) Run(cache *corpora.Corpora, language *config.ConfigLanguage, conf *config.Config) error {
	if len(c.Corpuse) == 0 {
		c.Corpuse = cache.SelectCorpuses()
	}
	be := conf.BatchEdit()
	for _, e := range c.Corpuse {
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
	if len(c.Corpuse) == 0 {
		c.Corpuse, _ = pterm.DefaultInteractiveMultiselect.
			WithOptions(language.Corpuses).
			Show()
	}
	be := conf.BatchEdit()
	for _, e := range c.Corpuse {
		be.Add(func(_ *config.Config) bool {
			return language.RemoveCorpuse(e)
		})
	}
	if _, err := be.Execute(); err != nil {
		slog.Error(err.Error())
	}
	return nil
}
