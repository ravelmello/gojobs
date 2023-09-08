package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ravelmello/gojobs/schemas"
	"net/http"
)

func UpdateJobHandler(context *gin.Context) {
	request := UpdateJob{}

	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("Error in validation %v", err.Error())
		sendError(context, http.StatusBadRequest, err.Error())
		return
	}

	id := context.Query("id")

	if id == "" {
		sendError(context,
			http.StatusBadRequest, validateParamFormat("id", "queryParam").Error())
		return
	}

	job := schemas.Job{}

	if err := db.First(&job, id).Error; err != nil {
		sendError(context, http.StatusNotFound, fmt.Sprintf("Id %s not found", id))
		return
	}

	if request.Role != "" {
		job.Role = request.Role
	}
	if request.Company != "" {
		job.Company = request.Company
	}
	if request.Location != "" {
		job.Location = request.Location
	}
	if request.Remote != nil {
		job.Remote = *request.Remote
	}
	if request.Link != "" {
		job.Link = request.Link
	}
	if request.Salary > 0 {
		job.Salary = request.Salary
	}

	if err := db.Save(&job).Error; err != nil {
		logger.ErrorF("Error at update job %v", err.Error())
		sendError(context, http.StatusInternalServerError, "Error in update job")
		return
	}

	sendSuccess(context, http.StatusAccepted, "Updated successfully", job)

}
