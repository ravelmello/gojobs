package router

import "github.com/gin-gonic/gin"

func InitializeRoutes() {
	router := gin.Default()

	initializeRoutes(router)

	err := router.Run(":9090")
	if err != nil {
		panic("Error in construction router")
	}
}
