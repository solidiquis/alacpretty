package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func TmpFile(fileContent, fileName string) (string, string) {
	content := []byte(fileContent)
	dir, err := ioutil.TempDir("./", "tmp")
	Must(err)

	tmpFilePath := filepath.Join(dir, fileName)
	err = ioutil.WriteFile(tmpFilePath, content, 0666)
	Must(err)

	return tmpFilePath, dir
}

func RemoveTmpDir(filePath string) {
	os.RemoveAll(filePath)
}

func ReadTmpToString(filePath string) string {
	text, err := ioutil.ReadFile(filePath)
	Must(err)

	return string(text)
}

func EndTest(code int) {
	os.Exit(code)
}

const ExampleYaml string = `
font:
  size: 9.0

  normal:
	family: Fira Code
	style: Regular

  bold:
	family: Fira Code
	style: Bold

  italic:
	family: Fira Code
	style: Bold Italic

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
