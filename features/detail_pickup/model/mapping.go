package model

import "recycle/features/detail_pickup/entity"

// Mapping dari Main ke Model
func MapMainToModel(mainData entity.Main) DetailPickup {
	return DetailPickup{
		PickupId:    mainData.PickupId,
		RubbishId:   mainData.RubbishId,
		ItemWeight:  mainData.ItemWeight,
		TotalPoints: mainData.TotalPoints,
	}
}

// Mapping dari Model ke Main
func MapModelToMain(mainData DetailPickup) entity.Main {
	return entity.Main{
		Id:          mainData.Id,
		PickupId:    mainData.PickupId,
		RubbishId:   mainData.RubbishId,
		ItemWeight:  mainData.ItemWeight,
		TotalPoints: mainData.TotalPoints,
		CreatedAt:   mainData.CreatedAt,
		UpdatedAt:   mainData.UpdatedAt,
	}
}

func ModelToMainMapping(dataModel []DetailPickup) []entity.Main {
	var mainList []entity.Main
	for _, value := range dataModel {
		mainList = append(mainList, MapModelToMain(value))
	}
	return mainList
}
