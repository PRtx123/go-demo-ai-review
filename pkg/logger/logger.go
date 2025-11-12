package logger

import (
	"fmt"
	"log"
)

func Infof(format string, args ...interface{}) {
	log.Printf("[INFO] "+format, args...)
}

func Errorf(format string, args ...interface{}) {
	log.Printf("[ERROR] "+format, args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf("[FATAL] "+format, args...)
}

func Debug(v ...interface{}) {
	fmt.Println(v...)
}
