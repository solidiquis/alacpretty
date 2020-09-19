package main

import (
	"io/ioutil"
	"os"
	"testing"
)

const exampleYamlPath string = "./assets/example.yml"

var content string

func init() {
	if _, err := os.Stat(exampleYamlPath); err != nil {
		must(err)
	}

	contentBytes, _ := ioutil.ReadFile(exampleYamlPath)
	content = string(contentBytes)

}

func TestCurrentTheme(t *testing.T) {
	subject := currentTheme
	targetValue := "Solarized Light"

	result := subject(&content)
	if result != targetValue {
		t.Errorf("\nExpected result: %s\nActual result: %s", targetValue, result)
	}
}

func TestCurrentOpacity(t *testing.T) {
	subject := currentOpacity
	targetValue := 0.9

	result := subject(&content)
	if result != targetValue {
		t.Errorf("\nExpected result: %.1f\nActual result: %.1f", targetValue, result)
	}
}
