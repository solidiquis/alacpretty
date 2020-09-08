package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
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
	fileContentBytes := []byte(*fileContent)

	newTheme := func(theme string) string {
		switch theme {
		case "solarized_light":
			return solarizedLight
		case "ayu_dark":
			return ayuDark
		default:
			return defaultTheme
		}
	}(theme)

	setTheme := func(settings []string, regex string) {
		for _, i := range settings {
			pattern := fmt.Sprintf("%s:%s", i, regex)
			re, _ := regexp.Compile(pattern)

			matchInNew := re.Match([]byte(newTheme))
			matchInOld := re.Match([]byte(fileContentBytes))

			if matchInOld && matchInNew {
				newSettings := re.Find([]byte(newTheme))
				newSettings = append(newSettings, '\n')
				*fileContent = re.ReplaceAllString(*fileContent, string(newSettings))
			} else {
				*fileContent = re.ReplaceAllString(*fileContent, string(""))
			}
		}
	}

	primarySettings := []string{
		"primary",
	}

	colorSettings := []string{
		"normal",
		"bright",
		"dim",
	}

	groundsTemplateString := "(?:\\n\\s+(background|foreground):.*){2}"
	colorTemplateString := "(?:\\n\\s+(black|red|green|yellow|blue|magenta|cyan|white):.*){8}"

	setTheme(primarySettings, groundsTemplateString)
	setTheme(colorSettings, colorTemplateString)
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
	theme := flag.String("th", "default", "theme")
	flag.Parse()

	content := readFileToString()
	if *fontSize > 0 {
		changeFontSize(&content, *fontSize)
	}
	if *opacity > 0.0 {
		changeOpacity(&content, *opacity)
	}
	if *theme != "" {
		changeTheme(&content, *theme)
	}

	applyChanges(content)
}
