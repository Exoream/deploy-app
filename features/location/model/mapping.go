package model

import "recycle/features/location/entity"

// Mapping dari Main ke Model
func MapMainToModel(mainData entity.Main) Location {
	return Location{
		City:       mainData.City,
		Subdistric: mainData.Subdistric,
		PostalCode: mainData.PostalCode,
		Longitude:  mainData.Longitude,
		Latitude:   mainData.Latitude,
	}
}

// Mapping dari Model ke Main
func MapModelToMain(mainData Location) entity.Main {
	return entity.Main{
		Id:         mainData.Id,
		City:       mainData.City,
		Subdistric: mainData.Subdistric,
		PostalCode: mainData.PostalCode,
		Longitude:  mainData.Longitude,
		Latitude:   mainData.Latitude,
		CreatedAt:  mainData.CreatedAt,
		UpdatedAt:  mainData.UpdatedAt,
	}
}

func ModelToMainMapping(dataModel []Location) []entity.Main {
	var mainList []entity.Main
	for _, value := range dataModel {
		mainList = append(mainList, MapModelToMain(value))
	}
	return mainList
}
