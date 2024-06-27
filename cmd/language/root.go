package language

import (
	"path/filepath"

	"github.com/8VIM/keyboard_layout_calculator/cmd/flags"
	"github.com/8VIM/keyboard_layout_calculator/cmd/language/corpuse"
	"github.com/8VIM/keyboard_layout_calculator/config"
	"github.com/8VIM/keyboard_layout_calculator/corpora"
	"github.com/alecthomas/kong"
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
	if err = cache.Load(); err != nil {
		return
	}
	if c.Language == "" {
		if language, err = conf.SelectLanguage(); err != nil {
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
