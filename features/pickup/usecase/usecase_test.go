package usecase

import (
	"errors"
	"mime/multipart"
	"net/textproto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"recycle/features/pickup/entity"
	"recycle/features/pickup/mocks"
)

func TestCreateValidPickup(t *testing.T) {
	pickupRepo := new(mocks.PickupDataInterface)
	usecase := NewPickupUsecase(pickupRepo)

	data := entity.Main{
		Address:    "minasaupa",
		Longitude:  "1231231",
		Latitude:   "421321",
		PickupDate: "2021-09-10",
		LocationId: "321321",
		Status:     "schedule",
	}

	pickupRepo.On("Create", mock.Anything, mock.AnythingOfType("*multipart.FileHeader")).Return(nil)
	emptyFileHeader := &multipart.FileHeader{}
	err := usecase.Create(data, emptyFileHeader)

	assert.NoError(t, err)
	pickupRepo.AssertCalled(t, "Create", data, emptyFileHeader)
}

func TestCreateInvalidPickup(t *testing.T) {
	pickupRepo := new(mocks.PickupDataInterface)
	usecase := NewPickupUsecase(pickupRepo)

	data := entity.Main{
		Address:    "minasaupa",
		Longitude:  "1231231",
		Latitude:   "421321",
		PickupDate: "2021-09-10",
		LocationId: "321321",
		Status:     "schedule",
	}
	image := &multipart.FileHeader{Size: 20 * 1024 * 1024}

	pickupRepo.On("Create", data, image).Return(nil)

	err := usecase.Create(data, image)
	expectedErr := errors.New("image file size should be less than 10 MB")
	assert.Equal(t, expectedErr, err)

	pickupRepo.AssertNotCalled(t, "Create")
}

func TestDeleteValidPickup(t *testing.T) {
	pickupRepo := new(mocks.PickupDataInterface)
	usecase := NewPickupUsecase(pickupRepo)

	id := "123"

	pickupRepo.On("GetById", mock.Anything).Return(entity.Main{}, nil)
	pickupRepo.On("DeleteById", id).Return(nil)

	err := usecase.DeleteById(id)
	assert.NoError(t, err)

	pickupRepo.AssertCalled(t, "DeleteById", id)
}

func TestDeleteInvalidPickup(t *testing.T) {
	pickupRepo := new(mocks.PickupDataInterface)
	usecase := NewPickupUsecase(pickupRepo)

	id := "123"

	pickupRepo.On("GetById", mock.Anything).Return(entity.Main{}, nil)
	pickupRepo.On("DeleteById", id).Return(nil)

	usecase.DeleteById(id)

	pickupRepo.AssertNotCalled(t, "DeleteById")
}

func TestDeleteNonExistentPickup(t *testing.T) {
	pickupRepo := new(mocks.PickupDataInterface)
	usecase := NewPickupUsecase(pickupRepo)

	id := "non-existent-id"

	pickupRepo.On("GetById", id).Return(entity.Main{}, errors.New("id not found"))
	err := usecase.DeleteById(id)

	expectedErr := errors.New("id not found")
	assert.Equal(t, expectedErr, err)

	pickupRepo.AssertNotCalled(t, "DeleteById")
}

func TestCreateErrorValidation(t *testing.T) {
	pickupRepo := new(mocks.PickupDataInterface)
	userUC := NewPickupUsecase(pickupRepo)

	testData := entity.Main{
		Address:    "",
		LocationId: "",
		PickupDate: "",
		Longitude:  "",
		Latitude:   "",
	}

	fileHeader := &multipart.FileHeader{
		Filename: "test.jpg",
		Size:     1000,
		Header:   make(textproto.MIMEHeader),
	}

	err := userUC.Create(testData, fileHeader)
	assert.NotNil(t, err)
}

func TestCreateErrorDateFormat(t *testing.T) {
	pickupRepo := new(mocks.PickupDataInterface)
	userUC := NewPickupUsecase(pickupRepo)

	testData := entity.Main{
		Address:    "minasaupa",
		Longitude:  "1231231",
		Latitude:   "421321",
		LocationId: "321321",
		Status:     "schedule",
		PickupDate: "2022/10/26",
	}

	fileHeader := &multipart.FileHeader{
		Filename: "test.jpg",
		Size:     1000,
		Header:   make(textproto.MIMEHeader),
	}

	err := userUC.Create(testData, fileHeader)

	expectedErr := errors.New("PickupDate should be in the format YYYY-MM-DD")
	assert.Equal(t, expectedErr, err)
}

func TestCreateErrorCreate(t *testing.T) {
	pickupRepo := new(mocks.PickupDataInterface)
	userUC := NewPickupUsecase(pickupRepo)

	testData := entity.Main{
		Address:    "minasaupa",
		Longitude:  "1231231",
		Latitude:   "421321",
		LocationId: "321321",
		PickupDate: "2021-09-10",
		Status:     "schedule",
	}

	fileHeader := &multipart.FileHeader{
		Filename: "test.jpg",
		Size:     1000,
		Header:   make(textproto.MIMEHeader),
	}

	expectedErr := errors.New("error while creating pickup")
	pickupRepo.On("Create", testData, fileHeader).Return(expectedErr)

	err := userUC.Create(testData, fileHeader)
	assert.Equal(t, expectedErr, err)
}

func TestDeleteRubbishMissingID(t *testing.T) {
	usecase := NewPickupUsecase(nil)

	id := ""

	err := usecase.DeleteById(id)
	assert.Error(t, err)

	assert.Equal(t, err.Error(), "id not found")
}

func TestDeleteCompletedPickup(t *testing.T) {
	pickupRepo := new(mocks.PickupDataInterface)
	userUC := NewPickupUsecase(pickupRepo)

	id := "123"

	pickup := entity.Main{
		Status: "done",
	}

	pickupRepo.On("GetById", id).Return(pickup, nil)
	err := userUC.DeleteById(id)

	expectedErr := errors.New("pickup is already completed and cannot be deleted")
	assert.Equal(t, expectedErr, err)
}

func TestDeleteError(t *testing.T) {
	pickupRepo := new(mocks.PickupDataInterface)
	userUC := NewPickupUsecase(pickupRepo)

	id := "123"

	pickupRepo.On("GetById", id).Return(entity.Main{}, nil)
	expectedErr := errors.New("delete error")
	pickupRepo.On("DeleteById", id).Return(expectedErr)

	err := userUC.DeleteById(id)

	assert.Equal(t, expectedErr, err)
}

func TestUpdateByIdSuccess(t *testing.T) {
	pickupRepo := new(mocks.PickupDataInterface)
	usecase := NewPickupUsecase(pickupRepo)

	id := "123"
	updatedData := entity.Main{
		Address:    "minasaupa",
		Longitude:  "1231231",
		Latitude:   "421321",
		LocationId: "321321",
		PickupDate: "2021-09-10",
	}

	nonCompletedData := entity.Main{
		Status: "scheduled",
	}
	pickupRepo.On("GetById", id).Return(nonCompletedData, nil)
	pickupRepo.On("UpdateById", id, updatedData, mock.Anything).Return(updatedData, nil)

	returnedData, err := usecase.UpdateById(id, updatedData, nil)

	assert.NoError(t, err)
	assert.Equal(t, updatedData, returnedData)
}

func TestUpdateByIdIdNotFound(t *testing.T) {
	pickupRepo := new(mocks.PickupDataInterface)
	usecase := NewPickupUsecase(pickupRepo)

	id := ""
	updatedData := entity.Main{}

	_, err := usecase.UpdateById(id, updatedData, nil)

	expectedErr := errors.New("id not found")
	assert.Equal(t, expectedErr, err)
}

func TestUpdateByIdImageSizeTooLarge(t *testing.T) {
	pickupRepo := new(mocks.PickupDataInterface)
	usecase := NewPickupUsecase(pickupRepo)

	id := "123"
	updatedData := entity.Main{}

	image := &multipart.FileHeader{
		Size: 11 * 1024 * 1024,
	}

	_, err := usecase.UpdateById(id, updatedData, image)

	expectedErr := errors.New("image file size should be less than 10 MB")
	assert.Equal(t, expectedErr, err)
}

func TestUpdateByIdError(t *testing.T) {
	pickupRepo := new(mocks.PickupDataInterface)
	usecase := NewPickupUsecase(pickupRepo)

	id := "123"
	updatedData := entity.Main{}
	image := &multipart.FileHeader{}

	expectedErr := errors.New("update error")
	pickupRepo.On("UpdateById", id, updatedData, image).Return(entity.Main{}, expectedErr)

	_, err := usecase.UpdateById(id, updatedData, image)

	assert.Equal(t, expectedErr, err)
}

func TestUpdateByIdPickupAlreadyCompleted(t *testing.T) {
    pickupRepo := new(mocks.PickupDataInterface)
    usecase := NewPickupUsecase(pickupRepo)

    id := "123"
    updatedData := entity.Main{} // Create an empty entity.Main for updated data
    image := &multipart.FileHeader{} // Mock a valid image header

	updatedMain := entity.Main{
        Status: "done",
    }

    // Mock the UpdateById method to return an error
    pickupRepo.On("UpdateById", id, updatedData, image).Return(updatedMain, nil)

    _, err := usecase.UpdateById(id, updatedData, image)

    expectedErr := errors.New("pickup is already completed and cannot be updated or deleted")
    assert.Equal(t, expectedErr, err)
}

func TestGetPickupByIDSuccess(t *testing.T) {
	pickupRepo := new(mocks.PickupDataInterface)
	usecase := NewPickupUsecase(pickupRepo)

	id := "123"

	pickupData := entity.Main{
		Id:   id,
		Address: "misaupa",
	}

	pickupRepo.On("GetById", id).Return(pickupData, nil)

	data, err := usecase.GetById(id)
	assert.NoError(t, err)
	assert.NotNil(t, data)

	pickupRepo.AssertCalled(t, "GetById", id)
}

func TestGetPickupByIDInvalidID(t *testing.T) {
	usecase := NewPickupUsecase(nil)

	id := ""

	_, err := usecase.GetById(id)

	assert.Error(t, err)
	assert.EqualError(t, err, "invalid id")
}

func TestFindAllPickupSuccess(t *testing.T) {
	pickupRepo := new(mocks.PickupDataInterface)
	usecase := NewPickupUsecase(pickupRepo)

	rubbishList := []entity.Main{
		{
			Address: "minasaupa",
			Longitude: "421421",
			Latitude:  "4213421",
		},
		{
			Address: "minasaupa",
			Longitude: "421421",
			Latitude:  "4213421",
		},
	}

	pickupRepo.On("FindAllPickup").Return(rubbishList, nil)

	rubbish, err := usecase.FindAllPickup()
	assert.NoError(t, err)
	assert.NotNil(t, rubbish)

	pickupRepo.AssertCalled(t, "FindAllPickup")
}

func TestFindAllPickupError(t *testing.T) {
	pickupRepo := new(mocks.PickupDataInterface)
	usecase := NewPickupUsecase(pickupRepo)

	expectedError := errors.New("error message")
	pickupRepo.On("FindAllPickup").Return(nil, expectedError)

	_, err := usecase.FindAllPickup()

	assert.Error(t, err)
	assert.EqualError(t, err, expectedError.Error())

	pickupRepo.AssertCalled(t, "FindAllPickup")
}

func TestGetByStatusValid(t *testing.T) {
    pickupRepo := new(mocks.PickupDataInterface)
    usecase := NewPickupUsecase(pickupRepo)

    status := "scheduled"
    fakeData := []entity.Main{}

    pickupRepo.On("GetByStatus", status).Return(fakeData, nil)

    result, err := usecase.GetByStatus(status)

    assert.NoError(t, err)
    assert.Equal(t, fakeData, result)
}

func TestGetByStatusInvalidStatus(t *testing.T) {
    pickupRepo := new(mocks.PickupDataInterface)
    usecase := NewPickupUsecase(pickupRepo)

    status := ""
    expectedErr := errors.New("status parameter is required")

    result, err := usecase.GetByStatus(status)

    assert.Equal(t, expectedErr, err)
    assert.Nil(t, result)
}







