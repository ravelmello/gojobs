package config

import (
	"fmt"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	db, err = InitializePgsql()

	if err != nil {
		return fmt.Errorf("error in initialization of Postgres %v", err)
	}
	return nil
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}

func GetDatabase() *gorm.DB {
	return db
}
