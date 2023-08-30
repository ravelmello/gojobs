package router

import "github.com/gin-gonic/gin"

func InitializeRoutes() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := router.Run(":9090")
	if err != nil {
		panic("Error in construction router")
	}
}
