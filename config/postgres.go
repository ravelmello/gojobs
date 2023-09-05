package config

import (
	"github.com/ravelmello/gojobs/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePgsql() (*gorm.DB, error) {
	logger := GetLogger("postgres")

	sc := "host=localhost user=postgres password=159753 dbname=gojobs port=5432 sslmode=disable"
	logger.Info("Initializing database")
	db, err := gorm.Open(postgres.Open(sc), &gorm.Config{})

	if err != nil {
		logger.ErrorF("database returns error: %v", err)
		return nil, err
	}

	logger.InfoF("Database connection succeeded: %v", db)

	err = db.AutoMigrate(&schemas.Job{})

	if err != nil {
		logger.ErrorF("Database auto migration error %v", err)
		return nil, err
	}

	return db, nil
}
