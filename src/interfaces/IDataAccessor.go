package interfaces

import (
	"github.com/abhishekb91/petstore-openapi3/src/api"
	"github.com/abhishekb91/petstore-openapi3/src/models"
)

type IDataAccessor interface {
	AddPet(pet *models.Pet) (*api.Pet, *api.Error)
	GetPets() ([]*api.Pet, *api.Error)
	GetPetById(petId int64) (*api.Pet, *api.Error)
	UpdatePet(petId int64, pet *models.Pet) *api.Error
	DeletePet(petId int64) *api.Error
}
