package main

import (
	"github.com/ravelmello/gojobs/config"
	"github.com/ravelmello/gojobs/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")
	err := config.Init()

	if err != nil {
		logger.ErrorF("Error in creation: %v", err)
		return
	}

	logger.Info("Initializing routes")
	router.InitializeRoutes()

}
