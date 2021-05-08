package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/abhishekb91/petstore-openapi3/src/api"
	models "github.com/abhishekb91/petstore-openapi3/src/model"
)

type IDataAccessor interface {
	AddPet(pet *models.Pet) (*api.Pet, *api.Error)
	GetPets() ([]*api.Pet, *api.Error)
}

type DataAccessor struct {
	db             *gorm.DB
	connectionInfo DBConnection
}

func NewDataAccessor() IDataAccessor {
	return &DataAccessor{}
}

func (da *DataAccessor) connect() error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	da.db = db

	// Migrate the schema
	da.db.AutoMigrate(&models.Pet{})
	da.db.AutoMigrate(&models.Owner{})

	return nil
}
