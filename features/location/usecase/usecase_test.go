package usecase

import (
	"errors"
	"recycle/features/location/entity"
	"recycle/features/location/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateLocation(t *testing.T) {
	locationRepo := new(mocks.LocationDataInterface)
	usecase := NewLocationUsecase(locationRepo)

	data := entity.Main{
		City:       "makassar",
		Subdistric: "rappocini",
		PostalCode: "90222",
		Longitude:  "21421412",
		Latitude:   "4124321",
	}

	expectedErr := errors.New("test error")
	locationRepo.On("Create", data).Return(expectedErr)

	err := usecase.Create(data)
	assert.Error(t, err)
	assert.ErrorIs(t, err, expectedErr)
}

func TestDeleteLocationById(t *testing.T) {
	locationRepo := new(mocks.LocationDataInterface)
	usecase := NewLocationUsecase(locationRepo)

	id := "123"

	expectedErr := errors.New("test error")
	locationRepo.On("DeleteById", id).Return(expectedErr)

	err := usecase.DeleteById(id)
	assert.Error(t, err)
	assert.ErrorIs(t, err, expectedErr)
}

func TestFindAllLocation(t *testing.T) {
	locationRepo := new(mocks.LocationDataInterface)
	usecase := NewLocationUsecase(locationRepo)

	expectedData := []entity.Main{}
	locationRepo.On("FindAllLocation").Return(expectedData, nil)

	result, err := usecase.FindAllLocation()
	assert.NoError(t, err)

	assert.Equal(t, expectedData, result)
}

func TestGetLocationById(t *testing.T) {
	locationRepo := new(mocks.LocationDataInterface)
	usecase := NewLocationUsecase(locationRepo)

	id := "123"

	expectedData := entity.Main{
		City:       "makassar",
		Subdistric: "rappocini",
		PostalCode: "421321",
		Latitude:   "421321",
		Longitude:  "41232131",
	}
	locationRepo.On("GetById", id).Return(expectedData, nil)

	result, err := usecase.GetById(id)
	assert.NoError(t, err)
	assert.Equal(t, expectedData, result)
}

func TestGetLocationByCity(t *testing.T) {
	locationRepo := new(mocks.LocationDataInterface)
	usecase := NewLocationUsecase(locationRepo)

	city := "Jakarta"

	expectedData := []entity.Main{
		{
			City:       "makassar",
			Subdistric: "rappocini",
			PostalCode: "421321",
			Latitude:   "421321",
			Longitude:  "41232131",
		},
		{
			City:       "jakarta",
			Subdistric: "kelapa gading",
			PostalCode: "421321",
			Latitude:   "421321",
			Longitude:  "41232131",
		},
	}
	locationRepo.On("GetByCity", city).Return(expectedData, nil)

	result, err := usecase.GetByCity(city)
	assert.NoError(t, err)
	assert.Equal(t, expectedData, result)
}

func TestUpdateLocationById(t *testing.T) {
	locationRepo := new(mocks.LocationDataInterface)
	usecase := NewLocationUsecase(locationRepo)

	id := "123"

	updatedData := entity.Main{}

	expectedData := entity.Main{}
	locationRepo.On("UpdateById", id, updatedData).Return(expectedData, nil)

	result, err := usecase.UpdateById(id, updatedData)
	assert.NoError(t, err)
	assert.Equal(t, expectedData, result)
}

func TestCreateErrorValidation(t *testing.T) {
	locationRepo := new(mocks.LocationDataInterface)
	userUC := NewLocationUsecase(locationRepo)

	testData := entity.Main{
		City:       "",
		Subdistric: "",
		PostalCode: "",
		Longitude:  "",
		Latitude:   "",
	}

	err := userUC.Create(testData)
	assert.NotNil(t, err)
}

func TestCreateLocationSuccess(t *testing.T) {
    locationRepo := new(mocks.LocationDataInterface)
    usecase := NewLocationUsecase(locationRepo)

    data := entity.Main{
        City:       "makassar",
        Subdistric: "rappocini",
        PostalCode: "90222",
        Longitude:  "21421412",
        Latitude:   "4124321",
    }

    locationRepo.On("Create", data).Return(nil)

    err := usecase.Create(data)
    assert.NoError(t, err)
}

func TestDeleteIdNotFound(t *testing.T) {
	usecase := NewLocationUsecase(nil)

	id := ""

	err := usecase.DeleteById(id)
	assert.Error(t, err)

	assert.Equal(t, err.Error(), "id not found")
}

func TestDeleteLocationByIdSuccess(t *testing.T) {
    locationRepo := new(mocks.LocationDataInterface)
    usecase := NewLocationUsecase(locationRepo)

    validID := "123"
    locationRepo.On("DeleteById", validID).Return(nil)

    err := usecase.DeleteById(validID)

    assert.NoError(t, err)
}

func TestFindAllLocationError(t *testing.T) {
    locationRepo := new(mocks.LocationDataInterface)
    usecase := NewLocationUsecase(locationRepo)

    expectedErr := errors.New("test error")
    locationRepo.On("FindAllLocation").Return(nil, expectedErr)

    _, err := usecase.FindAllLocation()

    assert.Error(t, err)
    assert.ErrorIs(t, err, expectedErr)
}

func TestGetIdNotFound(t *testing.T) {
	usecase := NewLocationUsecase(nil)

	id := ""

	_, err := usecase.GetById(id)
	assert.Error(t, err)

	assert.Equal(t, err.Error(), "invalid id")
}

func TestGetCityNotFound(t *testing.T) {
	usecase := NewLocationUsecase(nil)

	city := ""

	_, err := usecase.GetByCity(city)
	assert.Error(t, err)

	assert.Equal(t, err.Error(), "City parameter is required")
}

func TestUpdateByIdNotFound(t *testing.T) {
	usecase := NewLocationUsecase(nil)

	id := ""

	_, err := usecase.UpdateById(id, entity.Main{})
	assert.Error(t, err)

	assert.Equal(t, err.Error(), "id not found")
}

func TestUpdateByIdError(t *testing.T) {
    locationRepo := new(mocks.LocationDataInterface)
    usecase := NewLocationUsecase(locationRepo)

    id := "123"
    updatedData := entity.Main{
        City: "surabaya",
    }


    expectedErr := errors.New("error updating location")
    locationRepo.On("UpdateById", id, updatedData).Return(entity.Main{}, expectedErr)

    _, err := usecase.UpdateById(id, updatedData)
    assert.Error(t, err)
    assert.ErrorIs(t, err, expectedErr)

    locationRepo.AssertCalled(t, "UpdateById", id, updatedData)
}






