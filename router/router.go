package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ravelmello/gojobs/config"
)

var (
	logger *config.Logger
)

func InitializeRoutes() {
	logger = config.GetLogger("routes")
	router := gin.Default()

	initializeRoutes(router)

	logger.Info("Routes initialized")

	err := router.Run(":9090")
	if err != nil {
		panic("Error in construction router")
	}
}
