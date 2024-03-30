package cli

import (
	"fmt"
	"strings"

	"log/slog"

	"github.com/8VIM/keyboard_layout_calculator/bigrams"
	"github.com/8VIM/keyboard_layout_calculator/layout"
	"github.com/pterm/pterm"
)

func initLogger(verbose int) {
	var lvl pterm.LogLevel
	switch verbose {
	case 0:
		lvl = pterm.LogLevelWarn
	case 1:
		lvl = pterm.LogLevelInfo
	case 2:
		lvl = pterm.LogLevelDebug
	default:
		lvl = pterm.LogLevelTrace
	}
	handler := pterm.NewSlogHandler(pterm.DefaultLogger.WithLevel(lvl))
	logger := slog.New(handler)
	slog.SetDefault(logger)
}

func logError(err error, spinner *pterm.SpinnerPrinter) {
	slog.Error(err.Error())
	spinner.Fail()
}

func Execute(c *Config, args []string) {
	initLogger(c.Verbose)
	multi := pterm.DefaultMultiPrinter
	_, _ = multi.Start()
	defer func() { _, _ = multi.Stop() }()

	layoutsToProcess := make(map[string]struct{})

	for _, arg := range args {
		arg = strings.ToLower(arg)
		if arg == "*" {
			arg = "all"
		}
		layoutsToProcess[arg] = struct{}{}
	}

	if _, ok := layoutsToProcess["all"]; ok {
		layoutsToProcess = make(map[string]struct{})
	}

	r, err := bigrams.New(&multi, c.Parallelism, c.Force)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	spinner, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start("Loading corpuses")
	for _, layout := range c.Config.Layouts {
		if _, ok := layoutsToProcess[strings.ToLower(layout.Name)]; len(layoutsToProcess) != 0 && !ok {
			continue
		}

		for name := range layout.Languages {
			name = strings.ToLower(name)
			if language, ok := c.Config.Languages[name]; ok {
				r.Add(name, language)
			}
		}
	}
	ngramsByLanguage := r.Load()
	spinner.Info("Corpuses loaded")
	for _, la := range c.Config.Layouts {
		if _, ok := layoutsToProcess[strings.ToLower(la.Name)]; len(layoutsToProcess) != 0 && !ok {
			continue
		}
		spinner, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start(fmt.Sprintf("[%s] Generating layout", la.Name))
		l := layout.New(la.Name)
		ngram := bigrams.NewNGram()

		for name, lc := range la.Languages {
			name = strings.ToLower(name)
			if tmp, ok := ngramsByLanguage[name]; ok {
				ngram.Merge(tmp, lc.Weight)
			}
		}

		if err := l.AddFromNGram(ngram); err != nil {
			logError(err, spinner)
			continue
		}
		ngram.Graphivz(c.Output, l.Info.Name)
		if err := l.Save(c.Output); err != nil {
			logError(err, spinner)
			continue
		}

		spinner.Success(fmt.Sprintf("[%s] Generated layout", la.Name))
	}
}
