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
