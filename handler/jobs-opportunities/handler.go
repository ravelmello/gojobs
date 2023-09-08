package handler

import (
	"github.com/ravelmello/gojobs/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitHandler() {
	logger = config.GetLogger("Handler")
	db = config.GetDatabase()
}
