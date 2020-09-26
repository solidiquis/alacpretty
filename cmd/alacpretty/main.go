package main

import (
	"flag"

	alacui "github.com/solidiquis/alacpretty/internal/ui"
	"github.com/solidiquis/alacpretty/internal/yamlconf"
)

func main() {
	fontSize := flag.Int("fs", 0, "font size")
	opacity := flag.Float64("op", 0.0, "opacity")
	theme := flag.String("th", "", "theme")
	interactive_mode := flag.Bool("i", false, "interactive mode")
	flag.Parse()

	content := yamlconf.ReadFileToString()
	if *interactive_mode {
		alacui.WidgetsController(&content)
	} else {
		if *fontSize > 0 {
			yamlconf.ChangeFontSize(&content, *fontSize)
		}
		if *opacity > 0.0 {
			yamlconf.ChangeOpacity(&content, *opacity)
		}
		if *theme != "" {
			yamlconf.ChangeTheme(&content, *theme)
		}
		yamlconf.ApplyChanges(content)
	}
}
