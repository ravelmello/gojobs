package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ravelmello/gojobs/schemas"
	"net/http"
)

func ListJobsHandler(context *gin.Context) {
	var jobs []schemas.Job

	if err := db.Find(&jobs).Error; err != nil {
		sendError(context, http.StatusNotFound, "Error in list jobs")
		return
	}

	sendSuccess(context, http.StatusOK, "Find all Jobs", jobs)
}
