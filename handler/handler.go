package handler

import (
	"github.com/thiagoeu/GOstd/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLoggers("handler")
	db = config.GetSqlite()
}
