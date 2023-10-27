package model

import (
	// "recycle/features/pickup/controller"
	pick "recycle/features/pickup/entity"
	"recycle/features/pickup/model"
	"recycle/features/user/entity"
)

// Mapping dari Main ke Model
func MapMainToModel(mainData entity.Main) User {
	return User{
		Name:              mainData.Name,
		Email:             mainData.Email,
		Password:          mainData.Password,
		Gender:            mainData.Gender,
		Age:               mainData.Age,
		Address:           mainData.Address,
		SaldoPoints:       mainData.SaldoPoints,
		Role:              mainData.Role,
		IsVerified:        mainData.IsVerified,
		VerificationToken: mainData.VerificationToken,
	}
}

// Mapping dari Model ke Main
func MapModelToMain(mainData User) entity.Main {
	return entity.Main{
		Id:                mainData.Id,
		Name:              mainData.Name,
		Email:             mainData.Email,
		Password:          mainData.Password,
		Gender:            mainData.Gender,
		Age:               mainData.Age,
		Address:           mainData.Address,
		SaldoPoints:       mainData.SaldoPoints,
		Role:              mainData.Role,
		IsVerified:        mainData.IsVerified,
		VerificationToken: mainData.VerificationToken,
		CreatedAt:         mainData.CreatedAt,
		UpdatedAt:         mainData.UpdatedAt,
	}
}

func ModelToMainMapping(dataModel []User) []entity.Main {
	var mainList []entity.Main
	for _, value := range dataModel {
		main := MapModelToMain(value)

		// Ambil data Pickups dan map ke dalam struktur entity.Main
		var pickups []pick.Main
		for _, pickup := range value.Pickups {
			pickupMapped := model.MapModelToMain(pickup)
			pickups = append(pickups, pickupMapped)
		}
		main.Pickups = pickups

		mainList = append(mainList, main)
	}
	return mainList
}
