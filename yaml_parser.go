package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	fontRegex    = `size:\s*\d{2}\.0`
	opacityRegex = `background_opacity:\s*\d+.\d+`
	colorsRegex  = `\bcolors:.*(?:\n\s{2,}.+)+`
)

func currentTheme(fileContent *string) string {
	fileContentBytes := []byte(*fileContent)

	for k, v := range allThemes {
		match, _ := regexp.Match(v, fileContentBytes)
		if match {
			return k
		}
	}

	return "defaultTheme"
}

func currentOpacity(fileContent *string) float64 {
	regex, _ := regexp.Compile(opacityRegex)
	matchBytes := regex.Find([]byte(*fileContent))

	var matchString string
	if matchBytes != nil {
		matchString = string(matchBytes)
	}
	result := strings.Split(matchString, ":")
	opacityValueString := strings.Trim(result[len(result)-1], " ")

	opacityValue, err := strconv.ParseFloat(opacityValueString, 2)
	must(err)

	return opacityValue
}

func changeFontSize(fileContent *string, fontSize int) {
	newFontSize := fmt.Sprintf("size: %d.0", fontSize)

	regex, _ := regexp.Compile(fontRegex)
	*fileContent = regex.ReplaceAllString(*fileContent, newFontSize)
}

func changeOpacity(fileContent *string, opacity float64) {
	newOpacity := fmt.Sprintf("background_opacity: %.1f", opacity)

	regex, _ := regexp.Compile(opacityRegex)
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

	regex, _ := regexp.Compile(colorsRegex)
	*fileContent = regex.ReplaceAllString(*fileContent, newTheme)
}

func applyChanges(newContent string) {
	err := os.Truncate(alacrittyYamlPath, 0)
	must(err)

	file, err := os.OpenFile(alacrittyYamlPath, os.O_APPEND|os.O_WRONLY, 0644)
	must(err)
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(newContent)

	writer.Flush()
}
