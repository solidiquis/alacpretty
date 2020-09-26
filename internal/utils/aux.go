package utils

import (
	"log"
	"os"
)

var errorLog *log.Logger

func init() {
	errorLog = log.New(os.Stdout, "ERROR:\t", log.Ldate|log.Ltime)
}

func Must(err error) {
	if err != nil {
		errorLog.Fatalln(err)
	}
}
