package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/abhishekb91/petstore-openapi3/src/interfaces"
	"github.com/abhishekb91/petstore-openapi3/src/models"
)

const (
	RecordNotFound = "Record not found"
)

type dataAccessor struct {
	db             *gorm.DB
	connectionInfo models.DBConnection
}

func NewDataAccessor() interfaces.IDataAccessor {
	return &dataAccessor{}
}

func (da *dataAccessor) connect() error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	da.db = db

	// Migrate the schema
	// da.db.AutoMigrate(&models.Pet{})

	return nil
}
