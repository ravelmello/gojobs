package router

import (
	"github.com/gin-gonic/gin"
	handler "github.com/ravelmello/gojobs/handler/opportunities"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/job-opening", handler.ShowJobHandler)
		v1.POST("/job-opening", handler.CreateJobHandler)
		v1.PUT("/job-opening", handler.UpdateJobHandler)
		v1.DELETE("/job-opening", handler.DeleteJobHandler)
		v1.GET("/all-job-opening", handler.ListJobsHandler)

	}
}
