package database

import (
	"gorm.io/gorm"

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

func NewDataAccessor(dbConn *gorm.DB) interfaces.IDataAccessor {
	return &dataAccessor{db: dbConn}
}
