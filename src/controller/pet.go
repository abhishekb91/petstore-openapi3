package controller

import (
	"net/http"

	"github.com/abhishekb91/petstore-openapi3/src/api"
	"github.com/abhishekb91/petstore-openapi3/src/database"
	models "github.com/abhishekb91/petstore-openapi3/src/model"
	"github.com/labstack/echo/v4"
)

var (
	DataAccessor = database.NewDataAccessor()
)

type PetStoreImpl struct{}

func (*PetStoreImpl) GetPets(ctx echo.Context) error {
	resp, dbErr := DataAccessor.GetPets()
	if dbErr != nil {
		return sendPetstoreError(ctx, dbErr.Code, dbErr.Message)
	}
	ctx.JSON(http.StatusCreated, resp)
	return nil
}

func (*PetStoreImpl) AddPet(ctx echo.Context) error {
	var newPet api.PetRequest
	err := ctx.Bind(&newPet)
	if err != nil {
		return sendPetstoreError(ctx, http.StatusBadRequest, "Invalid format for NewPet")
	}

	pet := &models.Pet{
		Name:   newPet.Name,
		Status: *newPet.Status,
	}

	resp, dbErr := DataAccessor.AddPet(pet)

	if dbErr != nil {
		return sendPetstoreError(ctx, dbErr.Code, dbErr.Message)
	}

	ctx.JSON(http.StatusCreated, resp)
	if err != nil {
		// Something really bad happened, tell Echo that our handler failed
		return err
	}
	return nil
}

func (*PetStoreImpl) DeletePet(ctx echo.Context, petId int64) error {
	return nil
}

func (*PetStoreImpl) GetPetById(ctx echo.Context, petId int64) error {
	return nil
}

func (*PetStoreImpl) UpdatePetById(ctx echo.Context, petId int64) error {
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
