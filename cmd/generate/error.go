package generate

import (
	"log/slog"

	"github.com/pterm/pterm"
)

func logError(err error, spinner *pterm.SpinnerPrinter) {
	slog.Error(err.Error())
	spinner.Fail()
}
