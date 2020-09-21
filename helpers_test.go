package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func tmpFile(fileContent, fileName string) (string, string) {
	content := []byte(fileContent)
	dir, err := ioutil.TempDir("./", "tmp")
	must(err)

	tmpFilePath := filepath.Join(dir, fileName)
	err = ioutil.WriteFile(tmpFilePath, content, 0666)
	must(err)

	return tmpFilePath, dir
}

func removeTmpDir(filePath string) {
	os.RemoveAll(filePath)
}

func readTmpToString(filePath string) string {
	text, err := ioutil.ReadFile(filePath)
	must(err)

	return string(text)
}

func endTest(code int) {
	os.Exit(code)
}
