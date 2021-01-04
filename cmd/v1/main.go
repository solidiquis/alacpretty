package main

import (
	aui "github.com/solidiquis/alacpretty/internal/ui"
	"github.com/solidiquis/alacpretty/internal/yamlconf"
)

func main() {
	content := yamlconf.ReadFileToString()
	aui.WidgetsControllerV1(
		&content,
		aui.NewThemeShuffler(0, 0, 25, 10),
		aui.NewFontsizeAdjuster(26, 0, 51, 10),
		aui.NewOpacityGauge(0, 11, 51, 14),
	)
}
