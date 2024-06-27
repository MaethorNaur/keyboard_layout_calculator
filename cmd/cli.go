package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/8VIM/keyboard_layout_calculator/cmd/generate"
	"github.com/8VIM/keyboard_layout_calculator/cmd/language"
	"github.com/8VIM/keyboard_layout_calculator/config"
	"github.com/8VIM/keyboard_layout_calculator/utils"
	"github.com/alecthomas/kong"
	"github.com/pterm/pterm"
)

type CLI struct {
	Verbose  int                  `help:"Enable logging from 'ERROR' to 'DEBUG'" short:"v" type:"counter"`
	Config   string               `help:"Location of config file" short:"c" default:"~/.config/keyboard_layout_calculator.yaml" type:"path"`
	Generate generate.GenerateCmd `cmd:"" help:"Generate one or more layouts"`
	Language language.LanguageCmd `cmd:"" help:"Manage the available languages"`
}

func (c *CLI) AfterApply(conf *config.Config) error {
	initLogger(c.Verbose)
	return initLayoutConfig(c.Config, conf)
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

func initLayoutConfig(path string, c *config.Config) error {
	c.File = filepath.Clean(path)
	dir, file := filepath.Split(path)
	info, err := os.Stat(c.File)
	if file == "" || (err == nil && info.IsDir()) {
		return fmt.Errorf("%s should not be a directory", c.File)
	}
	if errors.Is(err, fs.ErrNotExist) {
		slog.Info(fmt.Sprintf("%s doesn't exists. Preparing directories for it", c.File))
		if err := utils.EnsureDirExists(filepath.Clean(dir)); err != nil {
			return err
		}
	}
	c.Load()
	return nil
}

func exitOnError(err error) {
	if err == nil {
		return
	}
	slog.Error(err.Error())
	os.Exit(1)
}

func Run() {
	homeDir, err := os.UserHomeDir()
	exitOnError(err)

	dir := filepath.Join(homeDir, ".config", "keyboard_layout_calculator", "cache")

	if err := utils.EnsureDirExists(dir); err != nil {
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
