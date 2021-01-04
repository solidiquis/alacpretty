package main

import (
	aui "github.com/solidiquis/alacpretty/internal/ui"
	"github.com/solidiquis/alacpretty/internal/yamlconf"
)

func main() {
	content := yamlconf.ReadFileToString()
	aui.WidgetsControllerV1(
		&content,
		aui.NewThemeShuffler(0, 0, 20, 20),
	)
}
