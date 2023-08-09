package config

import (
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	return nil
}

func GetLoggers(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
