package main

import (
	"github.com/pablohenriques/goopportunities/config"
	"github.com/pablohenriques/goopportunities/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Initializar()
}
