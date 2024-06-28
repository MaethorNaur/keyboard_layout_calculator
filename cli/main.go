package main

import (
	"log/slog"
	"os"

	"github.com/8VIM/keyboard_layout_calculator/cli/generate"
	"github.com/8VIM/keyboard_layout_calculator/cli/language"
	"github.com/8VIM/keyboard_layout_calculator/cli/layout"
	"github.com/8VIM/keyboard_layout_calculator/common"
	"github.com/8VIM/keyboard_layout_calculator/common/config"
	"github.com/alecthomas/kong"
	"github.com/pterm/pterm"
)

type CLI struct {
	Verbose  int                  `help:"Enable logging from 'ERROR' to 'DEBUG'" short:"v" type:"counter"`
	Config   string               `help:"Location of config file" short:"c" default:"~/.config/keyboard_layout_calculator.yaml" type:"path"`
	Layout   layout.LayoutCmd     `cmd:"" help:"Manage layouts"`
	Generate generate.GenerateCmd `cmd:"" help:"Generate one or more layouts"`
	Language language.LanguageCmd `cmd:"" help:"Manage the available languages"`
}

func (c *CLI) AfterApply(conf *config.Config) error {
	initLogger(c.Verbose)
	return conf.Init(c.Config)
}

func initLogger(verbose int) {
	lvl := pterm.LogLevelDisabled
	if verbose > 0 {
		verbosity := pterm.LogLevel(verbose)
		if verbosity > pterm.LogLevelWarn {
			verbosity = pterm.LogLevelWarn
		}
		lvl = pterm.LogLevelFatal - verbosity
	}

	handler := pterm.NewSlogHandler(pterm.DefaultLogger.WithLevel(lvl))
	slog.SetDefault(slog.New(handler))
}

func exitOnError(err error) {
	if err == nil {
		return
	}
	slog.Error(err.Error())
	os.Exit(1)
}

func main() {
	dir, err := common.InitCache()
	if err != nil {
		exitOnError(err)
	}

	ctx := kong.Parse(&CLI{},
		kong.Name("keyboard_layout_calculator"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: false,
		}),
		kong.Bind(config.New()),
		kong.Vars{"cacheDir": dir},
	)
	err = ctx.Run()

	ctx.FatalIfErrorf(err)
}
