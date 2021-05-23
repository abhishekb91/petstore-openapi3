package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/abhishekb91/petstore-openapi3/src/api"
	"github.com/abhishekb91/petstore-openapi3/src/mocks"
	"github.com/abhishekb91/petstore-openapi3/src/utils"
)

var (
	testPet1 = &api.Pet{
		Id: 1,
		PetRequest: api.PetRequest{
			Name:   "Pet1",
			Status: utils.StringPtr("available"),
		},
	}

	testPet2 = &api.Pet{
		Id: 2,
		PetRequest: api.PetRequest{
			Name:   "Pet2",
			Status: utils.StringPtr("available"),
		},
	}
)

func createTestContext(method string, body io.Reader) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, "/", body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	return ctx, rec
}

func Test_GetPets_Error_DatabaseError(t *testing.T) {
	// Setup
	ctx, rec := createTestContext(http.MethodGet, nil)
	mockDA := &mocks.IDataAccessor{}
	mockDA.On("GetPets").Return(nil, &api.Error{Code: 400, Message: "Test Message"})
	c := NewSvcController(mockDA)

	// Assertions
	if assert.NoError(t, c.GetPets(ctx)) {
		var result map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &result)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, float64(400), result["code"])
		assert.Equal(t, "Test Message", result["message"])
	}
}

func Test_GetPets_Success(t *testing.T) {
	// Setup
	ctx, rec := createTestContext(http.MethodGet, nil)
	mockDA := &mocks.IDataAccessor{}
	mockDA.On("GetPets").Return([]*api.Pet{testPet1, testPet2}, nil)
	c := NewSvcController(mockDA)

	// Assertions
	if assert.NoError(t, c.GetPets(ctx)) {
		var result []map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &result)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, 2, len(result))
	}
}

func Test_AddPet_Error_InvalidData(t *testing.T) {
	// Setup
	req := "{foo: bar}"
	ctx, rec := createTestContext(http.MethodPost, strings.NewReader(req))
	c := NewSvcController(&mocks.IDataAccessor{})

	if assert.NoError(t, c.AddPet(ctx)) {
		var result map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &result)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, float64(400), result["code"])
		assert.Equal(t, InvalidPetFormat, result["message"])
	}
}

func Test_AddPet_Error_DatabaseError(t *testing.T) {
	// Setup
	req := &api.PetRequest{
		Name:   "Pet 1",
		Status: utils.StringPtr("available"),
	}
	s, _ := json.Marshal(req)
	ctx, rec := createTestContext(http.MethodPost, bytes.NewReader(s))
	mockDA := &mocks.IDataAccessor{}
	mockDA.On("AddPet", mock.AnythingOfType("*models.Pet")).Return(nil, &api.Error{Code: 400, Message: "Test Message"})
	c := NewSvcController(mockDA)

	if assert.NoError(t, c.AddPet(ctx)) {
		var result map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &result)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, float64(400), result["code"])
		assert.Equal(t, "Test Message", result["message"])
	}
}

func Test_AddPet_Success(t *testing.T) {
	// Setup
	req := &api.PetRequest{
		Name:   "Pet 1",
		Status: utils.StringPtr("available"),
	}
	s, _ := json.Marshal(req)
	ctx, rec := createTestContext(http.MethodPost, bytes.NewReader(s))
	mockDA := &mocks.IDataAccessor{}
	mockDA.On("AddPet", mock.AnythingOfType("*models.Pet")).Return(testPet1, nil)
	c := NewSvcController(mockDA)

	// Assertions
	if assert.NoError(t, c.AddPet(ctx)) {
		var result map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &result)

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, float64(testPet1.Id), result["id"])
	}
}

func Test_DeletePet_Error_DatabaseError(t *testing.T) {
	// Setup
	ctx, rec := createTestContext(http.MethodDelete, nil)
	mockDA := &mocks.IDataAccessor{}
	mockDA.On("DeletePet", mock.AnythingOfType("int64")).Return(&api.Error{Code: 400, Message: "Test Message"})
	c := NewSvcController(mockDA)

	// Assertions
	if assert.NoError(t, c.DeletePet(ctx, 1)) {
		var result map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &result)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, float64(400), result["code"])
		assert.Equal(t, "Test Message", result["message"])
	}
}

func Test_DeletePet_Success(t *testing.T) {
	// Setup
	ctx, rec := createTestContext(http.MethodDelete, nil)
	mockDA := &mocks.IDataAccessor{}
	mockDA.On("DeletePet", mock.AnythingOfType("int64")).Return(nil)
	c := NewSvcController(mockDA)

	// Assertions
	if assert.NoError(t, c.DeletePet(ctx, 1)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}

func Test_GetPetById_Error_DatabaseError(t *testing.T) {
	// Setup
	ctx, rec := createTestContext(http.MethodGet, nil)
	mockDA := &mocks.IDataAccessor{}
	mockDA.On("GetPetById", mock.AnythingOfType("int64")).Return(nil, &api.Error{Code: 400, Message: "Test Message"})
	c := NewSvcController(mockDA)

	// Assertions
	if assert.NoError(t, c.GetPetById(ctx, 1)) {
		var result map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &result)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, float64(400), result["code"])
		assert.Equal(t, "Test Message", result["message"])
	}
}

func Test_GetPetById_Success(t *testing.T) {
	// Setup
	ctx, rec := createTestContext(http.MethodGet, nil)
	mockDA := &mocks.IDataAccessor{}
	mockDA.On("GetPetById", mock.AnythingOfType("int64")).Return(testPet1, nil)
	c := NewSvcController(mockDA)

	// Assertions
	if assert.NoError(t, c.GetPetById(ctx, 1)) {
		var result map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &result)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, float64(testPet1.Id), result["id"])
	}
}

func Test_UpdatePetById_Error_InvalidData(t *testing.T) {
	// Setup
	req := "{foo: bar}"
	ctx, rec := createTestContext(http.MethodPut, strings.NewReader(req))
	c := NewSvcController(&mocks.IDataAccessor{})

	if assert.NoError(t, c.UpdatePetById(ctx, 1)) {
		var result map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &result)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, float64(400), result["code"])
		assert.Equal(t, InvalidPetFormat, result["message"])
	}
}

func Test_UpdatePetById_Error_DatabaseError(t *testing.T) {
	// Setup
	req := &api.PetRequest{
		Name:   "Pet 2",
		Status: utils.StringPtr("available"),
	}
	s, _ := json.Marshal(req)
	ctx, rec := createTestContext(http.MethodPut, bytes.NewReader(s))
	mockDA := &mocks.IDataAccessor{}
	mockDA.On("UpdatePet", mock.AnythingOfType("int64"), mock.AnythingOfType("*models.Pet")).Return(&api.Error{Code: 400, Message: "Test Message"})
	c := NewSvcController(mockDA)

	if assert.NoError(t, c.UpdatePetById(ctx, 2)) {
		var result map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &result)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, float64(400), result["code"])
		assert.Equal(t, "Test Message", result["message"])
	}
}

func Test_UpdatePetById_Success(t *testing.T) {
	// Setup
	req := &api.PetRequest{
		Name:   "Pet 1",
		Status: utils.StringPtr("available"),
	}
	s, _ := json.Marshal(req)
	ctx, rec := createTestContext(http.MethodPost, bytes.NewReader(s))
	mockDA := &mocks.IDataAccessor{}
	mockDA.On("UpdatePet", mock.AnythingOfType("int64"), mock.AnythingOfType("*models.Pet")).Return(nil)
	c := NewSvcController(mockDA)

	// Assertions
	if assert.NoError(t, c.UpdatePetById(ctx, 2)) {
		var result map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &result)

		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}
