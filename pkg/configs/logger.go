package configs

import (
	"fmt"
	"log"
)

type ConfigLogger struct {
	Fatal func(v ...interface{})
	Error func(v ...interface{})
	Info  func(v ...interface{})
}

// Internal config loader, by default uses log package loggers, but can be changed
var logger = ConfigLogger{
	Fatal: log.Fatal,
	Error: log.Println,
	Info:  log.Println,
}

// Define the logger to use by this package
func SetConfigLogger(l ConfigLogger) {
	logger = l
}

func Fatalf(format string, v ...interface{}) {
	logger.Fatal(fmt.Sprintf(format, v...))
}

func Errorf(format string, v ...interface{}) {
	logger.Error(fmt.Sprintf(format, v...))
}

func Infof(format string, v ...interface{}) {
	logger.Info(fmt.Sprintf(format, v...))
}
