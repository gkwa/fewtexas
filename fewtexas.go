package fewtexas

import (
	"log/slog"

	"github.com/taylormonacelli/limespeed"
)

func Main() int {
	slog.Debug("initial", "value", limespeed.GetMyPackageVar())

	limespeed.SetMyPackageVar("newvalue")

	slog.Debug("initial", "value", limespeed.GetMyPackageVar())
	return 0
}
