package main

import (
	"github.com/thiagoeu/GOstd/config"
	"github.com/thiagoeu/GOstd/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLoggers("main")

	// initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	// initialize router
	router.Initialize()
}
