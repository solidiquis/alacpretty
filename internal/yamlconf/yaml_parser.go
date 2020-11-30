package yamlconf

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/solidiquis/alacpretty/internal/themes"
	"github.com/solidiquis/alacpretty/internal/utils"
)

const (
	colorsRegex   = `\bcolors:.*(?:\n\s{2,}.+)+`
	fontRegex     = `family:.*`
	fontSizeRegex = `size:\s*\d+\.\d+`
	opacityRegex  = `background_opacity:\s*\d+\.\d+`
)

func CurrentTheme(fileContent *string) string {
	fileContentBytes := []byte(*fileContent)

	for k, v := range themes.AllThemes {
		match, _ := regexp.Match(v, fileContentBytes)
		if match {
			return k
		}
	}

	return "DefaultTheme"
}

func CurrentOpacity(fileContent *string) float64 {
	regex, _ := regexp.Compile(opacityRegex)
	matchBytes := regex.Find([]byte(*fileContent))

	var matchString string
	if matchBytes != nil {
		matchString = string(matchBytes)
	}
	result := strings.Split(matchString, ":")
	opacityValueString := strings.Trim(result[len(result)-1], " ")

	opacityValue, err := strconv.ParseFloat(opacityValueString, 2)
	utils.Must(err)

	return opacityValue
}

func CurrentFontSize(fileContent *string) int {
	regex, _ := regexp.Compile(fontSizeRegex)
	matchBytes := regex.Find([]byte(*fileContent))
	matchString := string(matchBytes)
	regex, _ = regexp.Compile(`\d+`)
	matchBytes = regex.Find([]byte(matchString))
	matchString = string(matchBytes)
	fontSize, err := strconv.Atoi(matchString)
	utils.Must(err)
	return fontSize
}

func CurrentFont(fileContent *string) string {
	regex, _ := regexp.Compile(fontRegex)
	matchBytes := regex.FindAll([]byte(*fileContent), -1)

	fontPattern := `\s+.*`
	regex, _ = regexp.Compile(fontPattern)

	var currentFont string
	if len(matchBytes) > 0 {
		currentFont = string(matchBytes[0])
		currentFont = string(regex.Find([]byte(currentFont)))
		currentFont = strings.Trim(currentFont, " ")
		currentFont = strings.Split(currentFont, ".")[0]
	}

	return currentFont
}

func ChangeFontSize(fileContent *string, fontSize int) {
	newFontSize := fmt.Sprintf("size: %d.0", fontSize)

	regex, _ := regexp.Compile(fontSizeRegex)
	*fileContent = regex.ReplaceAllString(*fileContent, newFontSize)
}

func ChangeOpacity(fileContent *string, opacity float64) {
	newOpacity := fmt.Sprintf("background_opacity: %.1f", opacity)

	regex, _ := regexp.Compile(opacityRegex)
	*fileContent = regex.ReplaceAllString(*fileContent, newOpacity)
}

func ChangeTheme(fileContent *string, theme string) {
	newTheme := func(theme string) string {
		switch theme {
		case "argonaut":
			return themes.Argonaut
		case "ayu_dark", "Ayu Dark":
			return themes.AyuDark
		case "ayu_mirage", "Ayu Mirage":
			return themes.AyuMirage
		case "after_glow", "After Glow":
			return themes.AfterGlow
		case "base16_default_dark", "Base16 Default Dark":
			return themes.Base16DefaultDark
		case "blood_moon", "Blood Moon":
			return themes.BloodMoon
		case "breeze", "Breeze":
			return themes.Breeze
		case "campbell", "Campbell":
			return themes.Campbell
		case "challenger_deep", "Challenger Deep":
			return themes.ChallengerDeep
		case "cobalt_2", "Cobalt2":
			return themes.Cobalt2
		case "cyber_punk_neon", "Cyber Punk Neon":
			return themes.CyberPunkNeon
		case "darcula", "Darcula":
			return themes.Darcula
		case "doom_one", "Doom One":
			return themes.DoomOne
		case "dracula", "Dracula":
			return themes.Dracula
		case "falcon", "Falcon":
			return themes.Falcon
		case "flat_remix", "Flat Remix":
			return themes.FlatRemix
		case "gotham", "Gotham":
			return themes.Gotham
		case "solarized_light", "Solarized Light":
			return themes.SolarizedLight
		default:
			return themes.DefaultTheme
		}
	}(theme)
	newTheme = strings.Trim(newTheme, "\n")

	regex, _ := regexp.Compile(colorsRegex)
	*fileContent = regex.ReplaceAllString(*fileContent, newTheme)
}

func ChangeFont(fileContent *string, font string) {
	newFont := fmt.Sprintf("family: %s", font)
	regex, _ := regexp.Compile(fontRegex)

	*fileContent = regex.ReplaceAllString(*fileContent, newFont)
}

// Intentionally not a pointer to a string
func TrimYamlPrefix(fileContent, splicePoint string) string {
	var patternToTrim string
	switch splicePoint {
	case "theme":
		patternToTrim = `\bcolors:(?:.*\n)+`
	case "font":
		patternToTrim = `font:(?:.*\n)+`
	case "opacity":
		patternToTrim = `background_opacity:(?:.*\n)+`
	}
	regex, _ := regexp.Compile(patternToTrim)
	matchBytes := regex.Find([]byte(fileContent))

	return string(matchBytes)
}

func ApplyChanges(newContent string) {
	err := os.Truncate(alacrittyYamlPath, 0)
	utils.Must(err)

	file, err := os.OpenFile(alacrittyYamlPath, os.O_APPEND|os.O_WRONLY, 0644)
	utils.Must(err)
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(newContent)

	writer.Flush()
}
