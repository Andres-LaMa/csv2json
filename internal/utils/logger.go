package utils

import (
	"fmt"
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

func LogError(msg interface{}) {
	switch v := msg.(type) {
	case error:
		ErrorLogger.Println(v.Error())
	case string:
		ErrorLogger.Println(v)
	default:
		ErrorLogger.Println(fmt.Sprint(v))
	}
}
