package yamlconf

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/solidiquis/alacpretty/internal/utils"
)

var alacrittyYamlPath string

func init() {
	findAlacrittyYaml()
}

func findAlacrittyYaml() {
	homeDir, err := os.UserHomeDir()
	utils.Must(err)

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

	utils.Must(err)
}

func ReadFileToString() string {
	text, err := ioutil.ReadFile(alacrittyYamlPath)
	utils.Must(err)

	return string(text)
}
