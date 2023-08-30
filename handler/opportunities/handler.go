package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateJobHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"Message": "Open opportunities",
	})
}

func UpdateJobHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"Message": "Open opportunities",
	})
}

func DeleteJobHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"Message": "Open opportunities",
	})
}

func ListJobsHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"Message": "Open opportunities",
	})
}

func ShowJobHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"Message": "Open opportunities",
	})
}
