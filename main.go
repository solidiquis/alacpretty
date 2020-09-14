package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

const alacrittyConfig string = "/Users/benjamin/dotfiles/alacritty.yml"

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func readFileToString() string {
	_, err := os.Stat(alacrittyConfig)
	must(err)

	text, err := ioutil.ReadFile(alacrittyConfig)
	must(err)

	return string(text)
}

func changeFontSize(fileContent *string, fontSize int) {
	newFontSize := fmt.Sprintf("size: %d.0", fontSize)

	regex, _ := regexp.Compile("size:\\s*\\d{2}\\.0")
	*fileContent = regex.ReplaceAllString(*fileContent, newFontSize)
}

func changeOpacity(fileContent *string, opacity float64) {
	newOpacity := fmt.Sprintf("background_opacity: %.1f", opacity)

	regex, _ := regexp.Compile("background_opacity:\\s*\\d+.\\d+")
	*fileContent = regex.ReplaceAllString(*fileContent, newOpacity)
}

func changeTheme(fileContent *string, theme string) {
	newTheme := func(theme string) string {
		switch theme {
		case "argonaut":
			return argonaut
		case "ayu_dark", "Ayu Dark":
			return ayuDark
		case "ayu_mirage", "Ayu Mirage":
			return ayuMirage
		case "after_glow", "After Glow":
			return afterGlow
		case "base16_default_dark", "Base16 Default Dark":
			return base16DefaultDark
		case "blood_moon", "Blood Moon":
			return bloodMoon
		case "solarized_light", "Solarized Light":
			return solarizedLight
		default:
			return defaultTheme
		}
	}(theme)
	newTheme = strings.Trim(newTheme, "\n")
	alacrittyColors := "\\bcolors:.*(?:\\n\\s{2,}.+)+"
	regex, _ := regexp.Compile(alacrittyColors)
	*fileContent = regex.ReplaceAllString(*fileContent, newTheme)
}

func applyChanges(newContent string) {
	err := os.Truncate(alacrittyConfig, 0)
	must(err)

	file, err := os.OpenFile(alacrittyConfig, os.O_APPEND|os.O_WRONLY, 0644)
	must(err)
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(newContent)

	writer.Flush()
}

func main() {
	fontSize := flag.Int("fs", 0, "font size")
	opacity := flag.Float64("op", 0.0, "opacity")
	theme := flag.String("th", "", "theme")
	interactive_mode := flag.Bool("i", false, "interactive mode")
	flag.Parse()

	content := readFileToString()
	if *interactive_mode {
		shuffleThemes(&content)
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
