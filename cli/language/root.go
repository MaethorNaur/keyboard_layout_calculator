package language

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/8VIM/keyboard_layout_calculator/cli/flags"
	"github.com/8VIM/keyboard_layout_calculator/cli/language/corpuse"
	"github.com/8VIM/keyboard_layout_calculator/common/config"
	"github.com/8VIM/keyboard_layout_calculator/common/corpora"
	"github.com/alecthomas/kong"
	"github.com/pterm/pterm"
)

type LanguageCmd struct {
	flags.FilesRetrieverFlags
	Language string             `short:"l" help:"Language to manage"`
	Corpuse  corpuse.CorpuseCmd `cmd:"" help:"Manage corpuses for a language"`
}

func (c *LanguageCmd) AfterApply(conf *config.Config, ctx *kong.Context, vars kong.Vars) (err error) {
	var language *config.ConfigLanguage
	cache := &corpora.Corpora{
		Parallelism: c.Parallelism,
		Force:       c.Force, CacheFile: filepath.Join(vars["cacheDir"], corpora.CacheFile),
	}
	ctx.Bind(cache)
	multi := &pterm.DefaultMultiPrinter
	spinner := pterm.DefaultSpinner.WithWriter(multi.NewWriter())
	p := pterm.DefaultProgressbar.WithWriter(multi.NewWriter())
	if err = cache.Load(func() {
		_, _ = multi.Start()
		spinner, _ = spinner.Start("Listing from Corpora")
	},
		func(i int) {
			p, _ = p.WithTotal(i).WithTitle("Languages retrieved").Start()
			spinner.UpdateText("Retrieving corpuses from Corpora")
		},
		func() { p.Increment() },
		func() {
			spinner.Info("Corpuses retrieved")
			multi.Stop()
		},
	); err != nil {
		return
	}
	if c.Language == "" {
		if language, err = selectLanguage(conf); err != nil {
			return
		}
	} else {
		if language, err = conf.Language(c.Language); err != nil {
			return
		}
	}

	ctx.Bind(language)
	return
}

func selectLanguage(c *config.Config) (l *config.ConfigLanguage, err error) {
	languages := c.LanguagesSize()
	if languages == 0 {
		name, _ := pterm.DefaultInteractiveTextInput.Show("Language name")
		if name == "" {
			err = fmt.Errorf("require a language name")
			return
		}

		letters, _ := pterm.DefaultInteractiveTextInput.Show("Allowed characters")
		if letters == "" {
			err = fmt.Errorf("require allowed characters for %s", name)
			return
		}

		name = strings.ToLower(name)
		c.AddLanguage(name, letters).Save()
		return
	}

	var selected string
	selected, err = pterm.DefaultInteractiveSelect.WithOptions(c.Languages()).Show()
	if err == nil {
		l, _ = c.Language(selected)
		l.SetName(selected)
	}
	return

}
