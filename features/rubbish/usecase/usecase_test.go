package usecase

import (
	"errors"
	"mime/multipart"
	"recycle/features/rubbish/entity"
	"recycle/features/rubbish/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateRubbishSuccess(t *testing.T) {
	rubbishRepo := new(mocks.RubbishDataInterface)
	usecase := NewRubbishUsecase(rubbishRepo)

	data := entity.Main{
		Name:        "botol",
		TypeRubbish: "plastik",
		PointPerKg:  1000,
		Description: "botol plastik",
	}

	rubbishRepo.On("Create", mock.Anything, mock.AnythingOfType("*multipart.FileHeader")).Return(nil)

	emptyFileHeader := &multipart.FileHeader{}
	err := usecase.Create(data, emptyFileHeader)
	assert.NoError(t, err)

	rubbishRepo.AssertCalled(t, "Create", data, emptyFileHeader)
}

func TestCreateRubbishValidation(t *testing.T) {
	usecase := NewRubbishUsecase(nil)

	data := entity.Main{}
	errCreate := usecase.Create(data, nil)

	assert.Error(t, errCreate)
	assert.Contains(t, errCreate.Error(), "Key: 'Main.Name' Error:Field validation for 'Name' failed on the 'required' tag")
	assert.Contains(t, errCreate.Error(), "Key: 'Main.TypeRubbish' Error:Field validation for 'TypeRubbish' failed on the 'required' tag")
	assert.Contains(t, errCreate.Error(), "Key: 'Main.PointPerKg' Error:Field validation for 'PointPerKg' failed on the 'required' tag")
}

func TestCreateRubbishImageSizeExceeded(t *testing.T) {
	rubbishRepo := new(mocks.RubbishDataInterface)
	usecase := NewRubbishUsecase(rubbishRepo)

	data := entity.Main{
		Name:        "botol",
		TypeRubbish: "plastik",
		PointPerKg:  1000,
		Description: "botol plastik",
	}

	rubbishRepo.On("Create", data, mock.AnythingOfType("*multipart.FileHeader")).Return(nil)
	image := &multipart.FileHeader{
		Size: 11 * 1024 * 1024,
	}

	err := usecase.Create(data, image)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "image file size should be less than 10 MB")

	rubbishRepo.AssertNotCalled(t, "Create")
}

func TestDeleteRubbishSuccess(t *testing.T) {
	rubbishRepo := new(mocks.RubbishDataInterface)
	usecase := NewRubbishUsecase(rubbishRepo)

	id := "123"

	rubbishRepo.On("DeleteById", id).Return(nil)

	err := usecase.DeleteById(id)
	assert.NoError(t, err)

	rubbishRepo.AssertCalled(t, "DeleteById", id)
}

func TestDeleteRubbishMissingID(t *testing.T) {
	usecase := NewRubbishUsecase(nil)

	id := ""

	err := usecase.DeleteById(id)
	assert.Error(t, err)

	assert.Equal(t, err.Error(), "id not found")
}

func TestFindAllRubbishSuccess(t *testing.T) {
	rubbishRepo := new(mocks.RubbishDataInterface)
	usecase := NewRubbishUsecase(rubbishRepo)

	rubbishList := []entity.Main{
		{
			Name:        "botol",
			TypeRubbish: "plastik",
			PointPerKg:  1000,
		},
		{
			Name:        "buku",
			TypeRubbish: "kertas",
			PointPerKg:  1000,
		},
	}

	rubbishRepo.On("FindAllRubbish").Return(rubbishList, nil)

	rubbish, err := usecase.FindAllRubbish()
	assert.NoError(t, err)
	assert.NotNil(t, rubbish)

	rubbishRepo.AssertCalled(t, "FindAllRubbish")
}

func TestFindAllRubbishError(t *testing.T) {
	rubbishRepo := new(mocks.RubbishDataInterface)
	usecase := NewRubbishUsecase(rubbishRepo)

	expectedError := errors.New("error message")
	rubbishRepo.On("FindAllRubbish").Return(nil, expectedError)

	_, err := usecase.FindAllRubbish()

	assert.Error(t, err)
	assert.EqualError(t, err, expectedError.Error())

	rubbishRepo.AssertCalled(t, "FindAllRubbish")
}

func TestGetRubbishByIDSuccess(t *testing.T) {
	rubbishRepo := new(mocks.RubbishDataInterface)
	usecase := NewRubbishUsecase(rubbishRepo)

	id := "123"

	rubbishData := entity.Main{
		Id:   id,
		Name: "kertas",
	}

	rubbishRepo.On("GetById", id).Return(rubbishData, nil)

	data, err := usecase.GetById(id)
	assert.NoError(t, err)
	assert.NotNil(t, data)

	rubbishRepo.AssertCalled(t, "GetById", id)
}

func TestGetRubbishByIDInvalidID(t *testing.T) {
	usecase := NewRubbishUsecase(nil)

	id := ""

	_, err := usecase.GetById(id)

	assert.Error(t, err)
	assert.EqualError(t, err, "invalid id")
}

func TestUpdateRubbishSuccess(t *testing.T) {
	rubbishRepo := new(mocks.RubbishDataInterface)
	usecase := NewRubbishUsecase(rubbishRepo)

	id := "123"
	updatedData := entity.Main{
		Id:          id,
		Name:        "karton",
		TypeRubbish: "karton",
	}

	rubbishRepo.On("UpdateById", id, updatedData, mock.AnythingOfType("*multipart.FileHeader")).Return(updatedData, nil)

	data, err := usecase.UpdateById(id, updatedData, nil)
	assert.NoError(t, err)
	assert.NotNil(t, data)

	rubbishRepo.AssertCalled(t, "UpdateById", id, updatedData, mock.AnythingOfType("*multipart.FileHeader"))
}

func TestUpdateRubbishInvalidID(t *testing.T) {
	rubbishRepo := new(mocks.RubbishDataInterface)
	usecase := NewRubbishUsecase(rubbishRepo)

	id := ""
	updatedData := entity.Main{
		Id:          id,
		Name:        "air minum",
		TypeRubbish: "plastik",
	}

	expectedErr := errors.New("id not found")
	rubbishRepo.On("UpdateById", id, updatedData, mock.AnythingOfType("*multipart.FileHeader")).Return(entity.Main{}, expectedErr)

	_, err := usecase.UpdateById(id, updatedData, nil)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
}

func TestGetRubbishByTypeSuccess(t *testing.T) {
	rubbishRepo := new(mocks.RubbishDataInterface)
	usecase := NewRubbishUsecase(rubbishRepo)

	typeRubbish := "Plastic"

	rubbishList := []entity.Main{
		{
			Name:        "botol",
			TypeRubbish: "plastik",
			PointPerKg:  1000,
		},
		{
			Name:        "buku",
			TypeRubbish: "kertas",
			PointPerKg:  1000,
		},
	}
	rubbishRepo.On("GetByType", typeRubbish).Return(rubbishList, nil)

	data, err := usecase.GetByType(typeRubbish)
	assert.NoError(t, err)
	assert.NotNil(t, data)

	rubbishRepo.AssertCalled(t, "GetByType", typeRubbish)
}

func TestGetRubbishByTypeMissingType(t *testing.T) {
	usecase := NewRubbishUsecase(nil)

	typeRubbish := ""

	_, err := usecase.GetByType(typeRubbish)
	assert.Error(t, err)

	expectedErr := "typeRubbish parameter is required"
	assert.EqualError(t, err, expectedErr)
}

func TestCreateRubbishError(t *testing.T) {
	rubbishRepo := new(mocks.RubbishDataInterface)
	usecase := NewRubbishUsecase(rubbishRepo)

	data := entity.Main{
		Name:        "buku",
		TypeRubbish: "plastik",
		PointPerKg:  1000,
	}
	image := &multipart.FileHeader{
		Size: 5 * 1024 * 1024,
	}

	expectedErr := errors.New("error while creating rubbish")
	rubbishRepo.On("Create", data, image).Return(expectedErr)

	err := usecase.Create(data, image)
	assert.Error(t, err)
	assert.EqualError(t, err, expectedErr.Error())

	rubbishRepo.AssertCalled(t, "Create", data, image)
}

func TestDeleteRubbishByIdError(t *testing.T) {
	rubbishRepo := new(mocks.RubbishDataInterface)
	usecase := NewRubbishUsecase(rubbishRepo)

	id := "123"

	expectedErr := errors.New("error while deleting rubbish by ID")
	rubbishRepo.On("DeleteById", id).Return(expectedErr)

	err := usecase.DeleteById(id)
	assert.Error(t, err)
	assert.EqualError(t, err, expectedErr.Error())

	rubbishRepo.AssertCalled(t, "DeleteById", id)
}

func TestCheckImageSizeExceeded(t *testing.T) {
	usecase := NewRubbishUsecase(nil)

	image := &multipart.FileHeader{
		Size: 11 * 1024 * 1024,
	}

	_, err := usecase.UpdateById("123", entity.Main{}, image)
	assert.Error(t, err)

	expectedErr := errors.New("image file size should be less than 10 MB")
	assert.EqualError(t, err, expectedErr.Error())
}

func TestUpdateByIdError(t *testing.T) {
	rubbishRepo := new(mocks.RubbishDataInterface)
	usecase := NewRubbishUsecase(rubbishRepo)

	id := "123"
	updatedData := entity.Main{}
	image := &multipart.FileHeader{}

	expectedErr := errors.New("test error")
	rubbishRepo.On("UpdateById", id, updatedData, image).Return(entity.Main{}, expectedErr)

	result, err := usecase.UpdateById(id, updatedData, image)
	
	assert.Error(t, err)
	assert.ErrorIs(t, err, expectedErr)
	assert.Equal(t, entity.Main{}, result)
}
