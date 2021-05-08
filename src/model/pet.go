package models

import (
	"github.com/abhishekb91/petstore-openapi3/src/api"
	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	Name        string
	Status      string
	VersionId   int64
	EnumTest    string  `gorm:"column:enumtest;type:ENUM('one','two','three');not null"`
	Owners      []Owner `gorm:"foreignKey:PetId;references:id"`
	LatestOwner Owner   `gorm:"foreignKey:PetId,VersionId;references:id,VersionId"`
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

type Owner struct {
	VersionId int64
	PetId     int64
	Name      string
	Address   string
}

func (Owner) TableName() string {
	return "Owner"
}
