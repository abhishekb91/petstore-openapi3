package controllers

import "github.com/abhishekb91/petstore-openapi3/src/interfaces"

const (
	InvalidPetFormat = "Invalid Pet format"
)

func NewSvcController(da interfaces.IDataAccessor) *SvcController {
	return &SvcController{
		dataAccessor: da,
	}
}

type SvcController struct {
	dataAccessor interfaces.IDataAccessor
}
