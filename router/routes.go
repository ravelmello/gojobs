package router

import (
	"github.com/gin-gonic/gin"
	jobsHandler "github.com/ravelmello/gojobs/handler/jobs-opportunities"
)

func initializeRoutes(router *gin.Engine) {

	logger.Info("Initializing Handler")
	//handler.InitHandler()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/job-opening", jobsHandler.ShowJobHandler)
		v1.POST("/job-opening", jobsHandler.CreateJobHandler)
		v1.PUT("/job-opening", jobsHandler.UpdateJobHandler)
		v1.DELETE("/job-opening", jobsHandler.DeleteJobHandler)
		v1.GET("/all-job-opening", jobsHandler.ListJobsHandler)

	}
}
