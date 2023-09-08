package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ravelmello/gojobs/schemas"
	"net/http"
)

func DeleteJobHandler(context *gin.Context) {

	id := context.Query("id")

	if id == "" {
		sendError(context,
			http.StatusBadRequest,
			validateParamFormat("id", "queryParam").Error())
		return
	}

	job := schemas.Job{}

	if err := db.First(&job, id).Error; err != nil {
		sendError(context, http.StatusNotFound, fmt.Sprintf("Id %s not found", id))
		return
	}

	if err := db.Delete(&job).Error; err != nil {
		sendError(context, http.StatusInternalServerError, fmt.Sprintf("Error in deletion job with id %s", id))
		return
	}

	sendSuccess(context, http.StatusMovedPermanently, "Delete Job", job)

}
