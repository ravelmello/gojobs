package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ravelmello/gojobs/config"
	"github.com/ravelmello/gojobs/schemas"
	"gorm.io/gorm"
	"net/http"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func CreateJobHandler(context *gin.Context) {
	request := Jobs{}
	logger = config.GetLogger("Create Job Handler")
	db = config.GetDatabase()

	logger.InfoF("Starting")

	errorBuild := context.BindJSON(&request)
	if errorBuild != nil {
		return
	}

	err := request.Validate()
	if err != nil {
		logger.ErrorF("Error in request %v", err.Error())
		sendError(context, http.StatusBadRequest, err.Error())
		return
	}

	logger.InfoF("Request Received %+v", request)

	job := schemas.Job{
		Role:        request.Role,
		Remote:      *request.Remote,
		Description: request.Description,
		Link:        request.Link,
		Company:     request.Company,
		Location:    request.Location,
		Salary:      request.Salary,
	}

	if err := db.Create(&job).Error; err != nil {
		logger.ErrorF("Error in creation %v", err.Error())
		sendError(context, http.StatusInternalServerError, "error in open database")
		return
	}
	sendSuccess(context, http.StatusCreated, "Operation in Create Job Handler, successfully", job)
}
