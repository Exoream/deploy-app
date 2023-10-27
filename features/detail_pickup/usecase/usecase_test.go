package usecase

import (
	"errors"
	"recycle/features/detail_pickup/entity"
	pickup "recycle/features/pickup/entity"
	rub "recycle/features/rubbish/entity"

	detail "recycle/features/detail_pickup/mocks"
	pick "recycle/features/pickup/mocks"
	rubbish "recycle/features/rubbish/mocks"
	data "recycle/features/user/entity"
	user "recycle/features/user/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// func TestCreateDetailPickup(t *testing.T) {
// 	// Membuat mock objects
// 	detailPickupRepo := new(detail.DetailPickupDataInterface)
// 	rubbishRepo := new(rubbish.RubbishDataInterface)
// 	pickupRepo := new(pick.PickupDataInterface)
// 	userRepo := new(user.UserDataInterface)

// 	// Membuat instance use case
// 	usecase := NewDetailPickupUsecase(detailPickupRepo, rubbishRepo, pickupRepo, userRepo)

// 	// Konfigurasi data palsu yang akan digunakan dalam pengujian
// 	fakeData := []entity.Main{
// 		{
// 			PickupId:   "1",
// 			RubbishId:  "2",
// 			ItemWeight: 5.0,
// 			TotalPoints: 50,
// 		},
// 	}

// 	// Mock pemanggilan method yang diperlukan
// 	detailPickupRepo.On("BeginTransaction").Return(nil, nil)
// 	rubbishRepo.On("GetById", "2").Return(rub.Main{PointPerKg: 10}, nil)
// 	pickupRepo.On("GetById", "1").Return(pickup.Main{Status: "schedule", UserId: "123"}, nil)
// 	detailPickupRepo.On("Create", fakeData).Return(nil)
// 	pickupRepo.On("UpdateStatus", "1", "done").Return(nil)
// 	userRepo.On("GetById", "123").Return(data.Main{SaldoPoints: 100}, nil)
// 	userRepo.On("UpdateById", "123", mock.AnythingOfType("entity.Main")).Return(nil)

// 	// Memanggil fungsi yang diuji
// 	totalPoints, err := usecase.Create(fakeData)

// 	// Memeriksa hasil pengujian
// 	assert.NoError(t, err)
// 	assert.Equal(t, 50, totalPoints)
// 	userRepo.AssertExpectations(t) // Memastikan semua panggilan ke userRepo terpanggil
// }

// func TestCreateDetailPickupInvalidPickupStatus(t *testing.T) {
// 	detailPickupRepo := new(detail.DetailPickupDataInterface)
// 	rubbishRepo := new(rubbish.RubbishDataInterface)
// 	pickupRepo := new(pick.PickupDataInterface)
// 	userRepo := new(user.UserDataInterface)

// 	usecase := NewDetailPickupUsecase(detailPickupRepo, rubbishRepo, pickupRepo, userRepo)

// 	fakeData := []entity.Main{
// 		{
// 			PickupId:   "1",
// 			RubbishId:  "2",
// 			ItemWeight: 5.0,
// 		},
// 		// Tambahkan data sesuai kebutuhan kasus Anda
// 	}

// 	// Mock pemanggilan method GetById
// 	detailPickupRepo.On("BeginTransaction").Return(nil, nil)
// 	rubbishRepo.On("GetById", "2").Return(rub.Main{PointPerKg: 10}, nil)
// 	pickupRepo.On("GetById", "1").Return(pickup.Main{Status: "done"}, nil)

// 	totalPoints, err := usecase.Create(fakeData)

// 	expectedErr := errors.New("cannot create detail pickup for a pickup with 'done' status")
// 	assert.Equal(t, expectedErr, err)
// 	assert.Zero(t, totalPoints)
// }

func TestCreateDetailPickup_Validation(t *testing.T) {
	// Membuat mock objects
	detailPickupRepo := new(detail.DetailPickupDataInterface)
	rubbishRepo := new(rubbish.RubbishDataInterface)
	pickupRepo := new(pick.PickupDataInterface)
	userRepo := new(user.UserDataInterface)

	// Membuat instance use case
	usecase := NewDetailPickupUsecase(detailPickupRepo, rubbishRepo, pickupRepo, userRepo)

	// Konfigurasi data palsu yang akan digunakan dalam pengujian
	fakeData := []entity.Main{
		{
			PickupId:    "1",
			RubbishId:   "2",
			ItemWeight:  5.0,
			TotalPoints: 50,
		},
	}

	// Mock pemanggilan method yang diperlukan
	detailPickupRepo.On("BeginTransaction").Return(nil)
	rubbishRepo.On("GetById", "2").Return(rub.Main{PointPerKg: 10}, nil)
	pickupRepo.On("GetById", "1").Return(pickup.Main{Status: "scheduled", UserId: "123"}, nil)
	detailPickupRepo.On("Create", fakeData).Return(nil)
	pickupRepo.On("UpdateStatus", "1", "done").Return(nil)
	userRepo.On("GetById", "123").Return(data.Main{SaldoPoints: 100}, nil)

	// Configure the UpdateById mock
	userRepo.On("UpdateById", "123", mock.AnythingOfType("entity.Main")).Return(nil)

	// Create an invalid data for testing validation
	invalidData := []entity.Main{
		{
			// Missing required fields, should trigger a validation error
		},
	}

	// Memanggil fungsi yang diuji with invalid data
	totalPoints, err := usecase.Create(invalidData)

	// Memeriksa hasil pengujian
	assert.Error(t, err)
	assert.Equal(t, 0, totalPoints)
}

func TestTransactionError(t *testing.T) {
	// Membuat mock objek detailPickupRepo
	detailPickupRepo := new(detail.DetailPickupDataInterface)
	rubbishRepo := new(rubbish.RubbishDataInterface)
	pickupRepo := new(pick.PickupDataInterface)
	userRepo := new(user.UserDataInterface)

	// Membuat instance use case
	usecase := NewDetailPickupUsecase(detailPickupRepo, rubbishRepo, pickupRepo, userRepo)

	// Konfigurasi mock untuk memicu kesalahan saat memulai transaksi
	detailPickupRepo.On("BeginTransaction").Return(nil, errors.New("Transaction error"))

	// Data palsu yang valid
	validData := []entity.Main{
		// Data yang valid
	}

	// Memanggil fungsi yang diuji
	_, err := usecase.Create(validData)

	// Memeriksa hasil pengujian, harapannya adalah ada error
	assert.Error(t, err)
	detailPickupRepo.AssertExpectations(t)
}

func TestFindAllDetailPickup(t *testing.T) {
	// Membuat mock objek repositori
	detailPickupRepo := new(detail.DetailPickupDataInterface)

	// Membuat instance use case
	usecase := NewDetailPickupUsecase(detailPickupRepo, nil, nil, nil)

	// Data palsu yang akan diharapkan dikembalikan oleh repositori
	fakeDetailPickupData := []entity.Main{
		{
			PickupId:   "1",
			RubbishId:  "2",
			ItemWeight: 3.0,
		},
		{
			PickupId:   "2",
			RubbishId:  "3",
			ItemWeight: 2.5,
		},
	}

	// Konfigurasi pemanggilan FindAllDetailPickup pada objek mock detailPickupRepo
	detailPickupRepo.On("FindAllDetailPickup").Return(fakeDetailPickupData, nil)

	// Memanggil fungsi yang diuji
	detailPickup, err := usecase.FindAllDetailPickup()

	// Memeriksa hasil pengujian
	assert.NoError(t, err)
	assert.NotNil(t, detailPickup)
	assert.Len(t, detailPickup, len(fakeDetailPickupData))

	// Memeriksa apakah hasil pengujian sama dengan data palsu
	for i, expected := range fakeDetailPickupData {
		assert.Equal(t, expected.PickupId, detailPickup[i].PickupId)
		assert.Equal(t, expected.RubbishId, detailPickup[i].RubbishId)
		assert.Equal(t, expected.ItemWeight, detailPickup[i].ItemWeight)
	}
}

func TestFindAllDetailPickup_Error(t *testing.T) {
	detailPickupRepo := new(detail.DetailPickupDataInterface)
	usecase := NewDetailPickupUsecase(detailPickupRepo, nil, nil, nil)

	expectedError := errors.New("database error")

	detailPickupRepo.On("FindAllDetailPickup").Return(nil, expectedError)
	detailPickup, err := usecase.FindAllDetailPickup()

	assert.Error(t, err)
	assert.Nil(t, detailPickup)
	assert.EqualError(t, err, expectedError.Error())
}