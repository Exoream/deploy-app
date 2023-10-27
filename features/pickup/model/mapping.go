package model

import "recycle/features/pickup/entity"

// Mapping dari Main ke Model
func MapMainToModel(mainData entity.Main) Pickup {
	return Pickup{
		Address:    mainData.Address,
		Longitude:  mainData.Longitude,
		Latitude:   mainData.Latitude,
		PickupDate: mainData.PickupDate,
		Status:     mainData.Status,
		UserId:     mainData.UserId,
		LocationId: mainData.LocationId,
		ImageURL:   mainData.ImageURL,
	}
}

// Mapping dari Model ke Main
func MapModelToMain(mainData Pickup) entity.Main {
	return entity.Main{
		Id:         mainData.Id,
		Address:    mainData.Address,
		Longitude:  mainData.Longitude,
		Latitude:   mainData.Latitude,
		PickupDate: mainData.PickupDate,
		Status:     mainData.Status,
		UserId:     mainData.UserId,
		LocationId: mainData.LocationId,
		ImageURL:   mainData.ImageURL,
		CreatedAt:  mainData.CreatedAt,
		UpdatedAt:  mainData.UpdatedAt,
	}
}

func ModelToMainMapping(dataModel []Pickup) []entity.Main {
	var mainList []entity.Main
	for _, value := range dataModel {
		mainList = append(mainList, MapModelToMain(value))
	}
	return mainList
}
