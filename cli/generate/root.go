package generate

import (
	"fmt"
	"path/filepath"

	"github.com/8VIM/keyboard_layout_calculator/cli/flags"
	"github.com/8VIM/keyboard_layout_calculator/common/bigrams"
	"github.com/8VIM/keyboard_layout_calculator/common/config"
	"github.com/8VIM/keyboard_layout_calculator/common/layout"
	"github.com/alecthomas/kong"
	"github.com/pterm/pterm"
)

type GenerateCmd struct {
	flags.FilesRetrieverFlags
	Output  string   `help:"Output directory" short:"o" type:"existingdir" default:"."`
	Dot     bool     `help:"Generate a graphivz dot file for each layouts" short:"d"`
	Layouts []string `arg:"" optional:"" help:"Generate 'layout' defined in the config.  You can generate several layouts at the same time.  if 'layout' is '*' all defined layouts will be generated.  If no layout is given you can select from a prompt"`
}

func (c *GenerateCmd) Run(conf *config.Config, vars kong.Vars) error {
	multi := pterm.DefaultMultiPrinter
	_, _ = multi.Start()
	defer func() { _, _ = multi.Stop() }()

	layoutProcessing := conf.NewLayoutProcessing()
	if len(c.Layouts) == 0 {
		c.Layouts, _ = pterm.DefaultInteractiveMultiselect.
			WithOptions(layoutProcessing.Names()).
			// WithDefaultText("Select which layout(s) to generate").
			// WithKeyConfirm(keys.Enter).
			// WithKeySelect(keys.Space).
			Show()
	} else {
		for _, v := range c.Layouts {
			if v == "*" {
				c.Layouts = make([]string, 0)
				break
			}
		}
	}
	layoutProcessing.Filter(c.Layouts)
	globalSpinner, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start("Loading corpuses")
	runner := bigrams.New(
		c.Parallelism,
		c.Force,
		vars["cacheDir"],
		func() *bigrams.Callbacks {
			spinner := pterm.DefaultSpinner.WithWriter(multi.NewWriter())
			return bigrams.NewCallbacks(
				func(s string) {
					spinner, _ = spinner.Start(fmt.Sprintf("[%s corpuses] Loading", s))
				},
				func(action byte, language, filename, size string) {
					if action == bigrams.Downloading {
						spinner.UpdateText(fmt.Sprintf("[%s corpuses] Downloading: %s (%s)", language, filename, size))
					} else {
						spinner.UpdateText(fmt.Sprintf("[%s corpuses] Extracting: %s (%s)", language, filename, size))
					}
				},
				func(s string) {
					spinner.UpdateText(fmt.Sprintf("[%s] Computing bigrams", s))
				},
				func(s string, err error) {
					spinner.Fail(fmt.Sprintf("[%s corpuses] Failed: %s", s, err))
				},
				func(s string) {
					spinner.Success(fmt.Sprintf("[%s corpuses] Done", s))
				},
			)
		},
	)
	for _, layout := range layoutProcessing.Layouts() {
		for _, l := range layout {
			runner.Add(l.Language())
		}
	}
	ngramsByLanguage := runner.Load()
	globalSpinner.Info("Corpuses loaded")

	for name, languages := range layoutProcessing.Layouts() {
		spinner, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start(fmt.Sprintf("[%s] Generating layout", name))
		l := layout.New(name)
		ngram := bigrams.NewNGram()

		for _, language := range languages {
			if tmp, ok := ngramsByLanguage[language.Language().Name()]; ok {
				ngram.Merge(tmp, language.Options().Weight)
			}
		}

		if err := l.AddFromNGram(ngram); err != nil {
			logError(err, spinner)
			continue
		}

		if err := l.Save(c.Output); err != nil {
			logError(err, spinner)
			continue
		}

		if c.Dot {
			name := filepath.Join(c.Output, fmt.Sprintf("%s.dot", l.Info.Name))
			if err := ngram.Dot(name); err != nil {
				logError(err, spinner)
				continue
			}
		}

		spinner.Success(fmt.Sprintf("[%s] Generated layout", name))
	}
	return nil
}
