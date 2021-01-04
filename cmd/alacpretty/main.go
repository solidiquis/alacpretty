package main

import (
	aui "github.com/solidiquis/alacpretty/internal/ui"
	"github.com/solidiquis/alacpretty/internal/yamlconf"
)

func main() {
	content := yamlconf.ReadFileToString()
	row1 := []aui.UIWidget{
		aui.NewThemeShuffler(0, 0, 25, 10),
		aui.NewFontsizeAdjuster(26, 0, 51, 10),
	}
	row2 := []aui.UIWidget{
		aui.NewOpacityGauge(0, 11, 51, 14),
	}

	aui.WidgetsController(
		&content,
		row1,
		row2,
	)
}
