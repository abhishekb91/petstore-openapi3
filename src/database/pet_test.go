package database

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"log"
	"regexp"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	GormLogger "gorm.io/gorm/logger"

	"github.com/abhishekb91/petstore-openapi3/src/interfaces"
	"github.com/abhishekb91/petstore-openapi3/src/models"
)

var (
	testPet1 = &models.Pet{
		Name:   "Pet 1",
		Status: "available",
	}

	testPet2 = &models.Pet{
		Name:   "Pet 2",
		Status: "available",
	}

	testInsertPet1 = &models.Pet{
		Name:   "Pet 3",
		Status: "available",
	}
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func initTests() (interfaces.IDataAccessor, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()

	gdb, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: GormLogger.Default.LogMode(GormLogger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect database, Error: %v", err)
	}

	mockDa := NewDataAccessor(gdb)

	return mockDa, mock
}

func Test_GetPets_Error_DatabaseError(t *testing.T) {
	mockDa, mock := initTests()

	mock.ExpectQuery("SELECT").
		WillReturnError(errors.New(""))

	pets, err := mockDa.GetPets()
	assert.Nil(t, pets)
	assert.Equal(t, int64(500), err.Code)
	assert.Equal(t, "Failed to get pets", err.Message)
}

func Test_GetPets_Success(t *testing.T) {
	mockDa, mock := initTests()

	mock.ExpectQuery("SELECT").
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "status"}).
			AddRow(1, time.Now(), time.Now(), nil, testPet1.Name, testPet1.Status).
			AddRow(2, time.Now(), time.Now(), nil, testPet1.Name, testPet1.Status))

	pets, err := mockDa.GetPets()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(pets))
}

func Test_AddPet_Error_DatabaseError(t *testing.T) {
	mockDa, mock := initTests()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `Pet`")).
		WithArgs(AnyTime{}, AnyTime{}, nil, testPet1.Name, testPet1.Status).
		WillReturnError(errors.New(""))
	mock.ExpectCommit()
	mock.ExpectClose()

	pet, err := mockDa.AddPet(testPet1)
	assert.Nil(t, pet)
	assert.Equal(t, int64(500), err.Code)
	assert.Equal(t, fmt.Sprintf("Failed to create pet for %v", testPet1.Name), err.Message)
}

func Test_AddPet_Success(t *testing.T) {
	mockDa, mock := initTests()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `Pet`")).
		WithArgs(AnyTime{}, AnyTime{}, nil, testInsertPet1.Name, testInsertPet1.Status).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	mock.ExpectClose()

	pet, err := mockDa.AddPet(testInsertPet1)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), pet.Id)
}

func Test_DeletePet_Error_DatabaseError(t *testing.T) {
	mockDa, mock := initTests()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE `Pet`")).
		WithArgs(AnyTime{}, 1).
		WillReturnError(errors.New(""))
	mock.ExpectCommit()
	mock.ExpectClose()

	err := mockDa.DeletePet(1)
	assert.NotNil(t, err)
	assert.Equal(t, RecordNotFound, err.Message)
	assert.Equal(t, int64(500), err.Code)
}

func Test_DeletePet_Success(t *testing.T) {
	mockDa, mock := initTests()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE `Pet`")).
		WithArgs(AnyTime{}, 1).
		WillReturnResult(sqlmock.NewResult(1, 0))
	mock.ExpectCommit()
	mock.ExpectClose()

	err := mockDa.DeletePet(1)
	assert.Nil(t, err)
}

func Test_GetPetById_Error_DatabaseError(t *testing.T) {
	mockDa, mock := initTests()

	mock.ExpectQuery("SELECT").
		WithArgs(1).
		WillReturnError(errors.New(""))
	mock.ExpectCommit()
	mock.ExpectClose()

	pet, err := mockDa.GetPetById(1)
	assert.Nil(t, pet)
	assert.NotNil(t, err)
	assert.Equal(t, int64(500), err.Code)
	assert.Equal(t, RecordNotFound, err.Message)
}

func Test_GetPetById_Error_NoRecords(t *testing.T) {
	mockDa, mock := initTests()

	mock.ExpectQuery("SELECT").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "status"}))
	mock.ExpectCommit()
	mock.ExpectClose()

	pet, err := mockDa.GetPetById(1)
	assert.Nil(t, pet)
	assert.NotNil(t, err)
	assert.Equal(t, int64(404), err.Code)
	assert.Equal(t, RecordNotFound, err.Message)
}

func Test_GetPetById_Success(t *testing.T) {
	mockDa, mock := initTests()

	mock.ExpectQuery("SELECT").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "status"}).
			AddRow(1, time.Now(), time.Now(), nil, testPet1.Name, testPet1.Status))
	mock.ExpectCommit()
	mock.ExpectClose()

	pet, err := mockDa.GetPetById(1)
	assert.Nil(t, err)
	assert.NotNil(t, pet)
}

func Test_UpdatePet_Error_DatabaseError(t *testing.T) {
	mockDa, mock := initTests()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE `Pet`")).
		WithArgs(AnyTime{}, testPet2.Name, testPet2.Status, 2).
		WillReturnError(errors.New(""))
	mock.ExpectCommit()
	mock.ExpectClose()

	err := mockDa.UpdatePet(2, testPet2)
	assert.NotNil(t, err)
	assert.Equal(t, RecordNotFound, err.Message)
	assert.Equal(t, int64(500), err.Code)
}

func Test_UpdatePet_Success(t *testing.T) {
	mockDa, mock := initTests()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE `Pet`")).
		WithArgs(AnyTime{}, testPet2.Name, testPet2.Status, 2).
		WillReturnResult(sqlmock.NewResult(2, 0))
	mock.ExpectCommit()
	mock.ExpectClose()

	err := mockDa.UpdatePet(2, testPet2)
	assert.Nil(t, err)
}
