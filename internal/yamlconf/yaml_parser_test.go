package yamlconf

import (
	"testing"

	"github.com/solidiquis/alacpretty/internal/utils"
)

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

func TestCurrentFont(t *testing.T) {
	subject := CurrentFont
	targetValue := "Fira Code"

	result := subject(&content)
	if result != targetValue {
		t.Errorf("\nExpected result: %s\nActual result: %s", targetValue, result)
	}
}

func TestChangeFont(t *testing.T) {
	subject := ChangeFont
	targetValue := "ArialHB"

	subject(&content, targetValue)
	result := CurrentFont(&content)
	if result != targetValue {
		t.Errorf("\nExpected result: %s\nActual result: %s", targetValue, result)
	}
}

// TODO
//func TestTrimYamlPrefix(t *testing.T) {
//subject := TrimYamlPrefix
//subject(content)
//}

func TestMain(m *testing.M) {
	tmp, dir := utils.TmpFile(utils.ExampleYaml, "tmp.yml")
	content = utils.ReadTmpToString(tmp)
	code := m.Run()
	utils.RemoveTmpDir(dir)
	utils.EndTest(code)
}
