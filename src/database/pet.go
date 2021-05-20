package database

import (
	"github.com/abhishekb91/petstore-openapi3/src/api"
	"github.com/abhishekb91/petstore-openapi3/src/models"
)

// GetPets returns all pets from the DB
func (da *dataAccessor) GetPets() ([]*api.Pet, *api.Error) {
	err := da.connect()
	if err != nil {
		msg := err.Error()
		return nil, &api.Error{Message: msg, Code: 500}
	}

	var pets []models.Pet
	var response []*api.Pet

	if err := da.db.Find(&pets).Error; err != nil {
		msg := "Failed to get pets"
		return nil, &api.Error{Message: msg, Code: 500}
	}

	for _, pet := range pets {
		response = append(response, pet.ToModel())
	}

	return response, nil
}

// AddPet adds a new pet to the DB
func (da *dataAccessor) AddPet(pet *models.Pet) (*api.Pet, *api.Error) {
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

// DeletePet deletes a pet from the DB
func (da *dataAccessor) DeletePet(petId int64) *api.Error {
	err := da.connect()
	if err != nil {
		msg := err.Error()
		return &api.Error{Message: msg, Code: 500}
	}

	petDTO := &models.Pet{}
	petDTO.ID = uint(petId)

	if err := da.db.Delete(&petDTO).Error; err != nil {
		msg := RecordNotFound
		return &api.Error{Message: msg, Code: 500}
	}

	return nil
}

// GetPetById gets a pet by id from the DB
func (da *dataAccessor) GetPetById(petId int64) (*api.Pet, *api.Error) {
	err := da.connect()
	if err != nil {
		msg := err.Error()
		return nil, &api.Error{Message: msg, Code: 500}
	}

	var pet models.Pet

	if err := da.db.Find(&pet, petId).Error; err != nil {
		msg := RecordNotFound
		return nil, &api.Error{Message: msg, Code: 500}
	}

	if pet.ID == 0 {
		return nil, &api.Error{Message: RecordNotFound, Code: 404}
	}

	return pet.ToModel(), nil
}

// UpdatePetById updates pet in the DB
func (da *dataAccessor) UpdatePet(petId int64, pet *models.Pet) *api.Error {
	err := da.connect()
	if err != nil {
		msg := err.Error()
		return &api.Error{Message: msg, Code: 500}
	}

	petDTO := &models.Pet{}
	petDTO.ID = uint(petId)

	if err := da.db.Model(&petDTO).Updates(pet).Error; err != nil {
		msg := RecordNotFound
		return &api.Error{Message: msg, Code: 500}
	}

	return nil
}
