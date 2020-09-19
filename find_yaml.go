package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var alacrittyYamlPath string

func init() {
	findAlacrittyYaml()
}

func findAlacrittyYaml() {
	homeDir, err := os.UserHomeDir()
	must(err)

	// https://github.com/alacritty/alacritty
	paths := []string{
		".config/alacritty/alacritty.yml",
		".alacritty.yml",
		"alacritty/.alacritty.yml",
	}

	var tmp string
	for _, path := range paths {
		tmp = fmt.Sprintf("%s/%s", homeDir, path)
		_, err := os.Stat(tmp)
		if err == nil {
			alacrittyYamlPath = tmp
			return
		}
	}

	must(err)
}

func readFileToString() string {
	text, err := ioutil.ReadFile(alacrittyYamlPath)
	must(err)

	return string(text)
}
