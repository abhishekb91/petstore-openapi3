package models

import (
	"github.com/abhishekb91/petstore-openapi3/src/api"
	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	Name   string
	Status string
}

func (Pet) TableName() string {
	return "Pet"
}

func (db *Pet) ToModel() *api.Pet {
	var pet = &api.Pet{}
	pet.Id = int64(db.ID)
	pet.Name = db.Name
	pet.Status = &db.Status
	return pet
}

func PetModelToDatabaseObject(petModel api.PetRequest) *Pet {
	var pet = &Pet{}
	pet.Name = petModel.Name
	pet.Status = *petModel.Status
	return pet
}
