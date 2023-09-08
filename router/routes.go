package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/ravelmello/gojobs/docs"
	jobsHandler "github.com/ravelmello/gojobs/handler/jobs-opportunities"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {

	jobsHandler.InitHandler()
	logger.Info("Initializing Handler")
	//handler.InitHandler()
	basePath := "/api/v1"

	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		v1.GET("/job-opening", jobsHandler.ShowJobHandler)
		v1.POST("/job-opening", jobsHandler.CreateJobHandler)
		v1.PUT("/job-opening", jobsHandler.UpdateJobHandler)
		v1.DELETE("/job-opening", jobsHandler.DeleteJobHandler)
		v1.GET("/all-job-opening", jobsHandler.ListJobsHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
