package main

import (
	"fmt"

	"github.com/thiagoeu/GOstd/config"
	"github.com/thiagoeu/GOstd/router"
)

func main() {
	// initialize configs
	err := config.Init()
	if err != nil {
		fmt.Print(err)
		return
	}

	// initialize router
	router.Initialize()
}
