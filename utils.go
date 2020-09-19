package main

import (
	"log"
	"os"
)

var errorLog *log.Logger

func init() {
	errorLog = log.New(os.Stdout, "ERROR:\t", log.Ldate|log.Ltime)
}

func must(err error) {
	if err != nil {
		errorLog.Fatalln(err)
	}
}
