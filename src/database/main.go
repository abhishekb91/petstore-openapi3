package database

import (
	"gorm.io/gorm"

	"github.com/abhishekb91/petstore-openapi3/src/interfaces"
)

const (
	RecordNotFound = "Record not found"
)

type dataAccessor struct {
	db *gorm.DB
}

func NewDataAccessor(dbConn *gorm.DB) interfaces.IDataAccessor {
	return &dataAccessor{db: dbConn}
}
