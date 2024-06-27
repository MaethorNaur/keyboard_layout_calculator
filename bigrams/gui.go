package bigrams

import (
	"github.com/8VIM/keyboard_layout_calculator/config"
	"github.com/pterm/pterm"
)

func LoadBigrams(l *config.LayoutProcessing, multi *pterm.MultiPrinter, parallelism int, force bool, cacheDir string) (result map[string]*NGram, err error) {
	r := New(multi, parallelism, force, cacheDir)
	spinner, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start("Loading corpuses")
	for _, layout := range l.Layouts() {
		for _, l := range layout {
			r.Add(l.Language())
		}
	}
	result = r.Load()
	spinner.Info("Corpuses loaded")
	return
}
