package controllers

import (
	"net/http"

	"github.com/abhishekb91/petstore-openapi3/src/api"
	"github.com/abhishekb91/petstore-openapi3/src/interfaces"
	"github.com/abhishekb91/petstore-openapi3/src/models"
	"github.com/labstack/echo/v4"
)

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

// GetPets returns all pets from the DB
func (c *SvcController) GetPets(ctx echo.Context) error {
	resp, dbErr := c.dataAccessor.GetPets()
	if dbErr != nil {
		return sendPetstoreError(ctx, dbErr.Code, dbErr.Message)
	}
	ctx.JSON(http.StatusOK, resp)
	return nil
}

// AddPet adds a new pet to the DB
func (c *SvcController) AddPet(ctx echo.Context) error {
	var newPet api.PetRequest
	err := ctx.Bind(&newPet)
	if err != nil {
		return sendPetstoreError(ctx, http.StatusBadRequest, InvalidPetFormat)
	}

	pet := &models.Pet{
		Name:   newPet.Name,
		Status: *newPet.Status,
	}

	resp, dbErr := c.dataAccessor.AddPet(pet)

	if dbErr != nil {
		return sendPetstoreError(ctx, dbErr.Code, dbErr.Message)
	}

	ctx.JSON(http.StatusCreated, resp)
	return nil
}

// DeletePet deletes a pet from the DB
func (c *SvcController) DeletePet(ctx echo.Context, petId int64) error {
	dbErr := c.dataAccessor.DeletePet(petId)
	if dbErr != nil {
		return sendPetstoreError(ctx, dbErr.Code, dbErr.Message)
	}
	ctx.NoContent(http.StatusNoContent)
	return nil
}

// GetPetById gets a pet by id from the DB
func (c *SvcController) GetPetById(ctx echo.Context, petId int64) error {
	resp, dbErr := c.dataAccessor.GetPetById(petId)
	if dbErr != nil {
		return sendPetstoreError(ctx, dbErr.Code, dbErr.Message)
	}
	ctx.JSON(http.StatusOK, resp)
	return nil
}

// UpdatePetById updates pet in the DB
func (c *SvcController) UpdatePetById(ctx echo.Context, petId int64) error {
	var updatePet api.PetRequest
	err := ctx.Bind(&updatePet)
	if err != nil {
		return sendPetstoreError(ctx, http.StatusBadRequest, InvalidPetFormat)
	}

	pet := models.PetModelToDatabaseObject(updatePet)

	dbErr := c.dataAccessor.UpdatePet(petId, pet)

	if dbErr != nil {
		return sendPetstoreError(ctx, dbErr.Code, dbErr.Message)
	}

	ctx.NoContent(http.StatusNoContent)
	return nil
}

// This function wraps sending of an error in the Error format, and
// handling the failure to marshal that.
func sendPetstoreError(ctx echo.Context, code int64, message string) error {
	petErr := api.Error{
		Code:    code,
		Message: message,
	}
	err := ctx.JSON(int(code), petErr)
	return err
}
