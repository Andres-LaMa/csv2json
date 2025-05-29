package utils

import (
	"log"
	"os"
)

var (
	InfoLogger  = log.New(os.Stdout, "[INFO] ", log.LstdFlags)
	ErrorLogger = log.New(os.Stderr, "[ERROR] ", log.LstdFlags)
)

func LogInfo(msg string) {
	InfoLogger.Println(msg)
}

func LogError(err error) {
	ErrorLogger.Println(err)
}
