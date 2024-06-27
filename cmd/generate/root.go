package generate

import (
	"fmt"
	"path/filepath"

	"github.com/8VIM/keyboard_layout_calculator/bigrams"
	"github.com/8VIM/keyboard_layout_calculator/cmd/flags"
	"github.com/8VIM/keyboard_layout_calculator/config"
	"github.com/8VIM/keyboard_layout_calculator/layout"
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

	layoutProcessing := conf.SelectLayouts(c.Layouts)
	ngramsByLanguage, err := bigrams.LoadBigrams(layoutProcessing, &multi, c.Parallelism, c.Force, vars["cacheDir"])

	if err != nil {
		return err
	}

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
