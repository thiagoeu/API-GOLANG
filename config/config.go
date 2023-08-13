package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// initialize sqlite
	db, err = InitializeSqlite()
	if err != nil {
		return fmt.Errorf("error initialize sqlite: %v", err)
	}

	return nil
}

func GetSqlite() *gorm.DB {
	return db
}

func GetLoggers(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
