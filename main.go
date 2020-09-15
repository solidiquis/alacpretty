package main

import (
	"flag"
)

func main() {
	fontSize := flag.Int("fs", 0, "font size")
	opacity := flag.Float64("op", 0.0, "opacity")
	theme := flag.String("th", "", "theme")
	interactive_mode := flag.Bool("i", false, "interactive mode")
	flag.Parse()

	content := readFileToString()
	if *interactive_mode {
		themeShuffler(&content)
	} else {
		if *fontSize > 0 {
			changeFontSize(&content, *fontSize)
		}
		if *opacity > 0.0 {
			changeOpacity(&content, *opacity)
		}
		if *theme != "" {
			changeTheme(&content, *theme)
		}
	}

	applyChanges(content)
}
