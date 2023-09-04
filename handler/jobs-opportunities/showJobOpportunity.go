package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowJobHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"Message": "Open jobs-opportunities",
	})
}
