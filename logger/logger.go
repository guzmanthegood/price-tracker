package logger

import (
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func init() {
	logger.Formatter = &logrus.TextFormatter{ForceColors: true}
	logger.SetLevel(logrus.InfoLevel)
}

// Info wrap
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Debug wrap
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Warn wrap
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Error wrap
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Fatal wrap
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Panic wrap
func Panic(args ...interface{}) {
	logger.Panic(args...)
}
