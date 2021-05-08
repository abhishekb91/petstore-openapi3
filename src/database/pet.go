package database

import (
	"github.com/abhishekb91/petstore-openapi3/src/api"
	models "github.com/abhishekb91/petstore-openapi3/src/model"
)

func (da *DataAccessor) AddPet(pet *models.Pet) (*api.Pet, *api.Error) {
	err := da.connect()
	if err != nil {
		msg := err.Error()
		return nil, &api.Error{Message: msg, Code: 500}
	}

	if err := da.db.Create(&pet).Error; err != nil {
		msg := "Failed to create pet for " + pet.Name
		return nil, &api.Error{Message: msg, Code: 500}
	}
	return pet.ToModel(), nil
}

func (da *DataAccessor) GetPets() ([]*api.Pet, *api.Error) {
	err := da.connect()
	if err != nil {
		msg := err.Error()
		return nil, &api.Error{Message: msg, Code: 500}
	}

	var pets []models.Pet

	if err := da.db.Preload("LatestOwner").Find(&pets).Error; err != nil {
		msg := "Failed to get pets"
		return nil, &api.Error{Message: msg, Code: 500}
	}

	return nil, nil
}
