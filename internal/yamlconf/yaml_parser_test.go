package yamlconf

import (
	"testing"

	"github.com/solidiquis/alacpretty/internal/utils"
)

const exampleYaml string = `
font:
  size: 9.0

colors:

  primary:
    background: '#fdf6e3'
    foreground: '#657b83'

  normal:
    black:   '#073642'
    red:     '#dc322f'
    green:   '#859900'
    yellow:  '#b58900'
    blue:    '#268bd2'
    magenta: '#d33682'
    cyan:    '#2aa198'
    white:   '#eee8d5'

  bright:
    black:   '#002b36'
    red:     '#cb4b16'
    green:   '#586e75'
    yellow:  '#657b83'
    blue:    '#839496'
    magenta: '#6c71c4'
    cyan:    '#93a1a1'
    white:   '#fdf6e3'

background_opacity: 0.9
`

var content string

func TestCurrentTheme(t *testing.T) {
	subject := CurrentTheme
	targetValue := "Solarized Light"

	result := subject(&content)
	if result != targetValue {
		t.Errorf("\nExpected result: %s\nActual result: %s", targetValue, result)
	}
}

func TestCurrentOpacity(t *testing.T) {
	subject := CurrentOpacity
	targetValue := 0.9

	result := subject(&content)
	if result != targetValue {
		t.Errorf("\nExpected result: %.1f\nActual result: %.1f", targetValue, result)
	}
}

func TestCurrentFontSize(t *testing.T) {
	subject := CurrentFontSize
	targetValue := 9

	result := subject(&content)
	if result != targetValue {
		t.Errorf("\nExpected result: %v\nActual result: %v", targetValue, result)
	}
}

func TestMain(m *testing.M) {
	tmp, dir := utils.TmpFile(exampleYaml, "tmp.yml")
	content = utils.ReadTmpToString(tmp)
	code := m.Run()
	utils.RemoveTmpDir(dir)
	utils.EndTest(code)
}
