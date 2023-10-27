package model

import "recycle/features/rubbish/entity"

// Mapping dari Main ke Model
func MapMainToModel(mainData entity.Main) Rubbish {
	return Rubbish{
		Name:        mainData.Name,
		TypeRubbish: mainData.TypeRubbish,
		PointPerKg:  mainData.PointPerKg,
		Description: mainData.Description,
		ImageURL:    mainData.ImageURL,
	}
}

// Mapping dari Model ke Main
func MapModelToMain(mainData Rubbish) entity.Main {
	return entity.Main{
		Id:          mainData.Id,
		Name:        mainData.Name,
		TypeRubbish: mainData.TypeRubbish,
		PointPerKg:  mainData.PointPerKg,
		Description: mainData.Description,
		ImageURL:    mainData.ImageURL,
		CreatedAt:   mainData.CreatedAt,
		UpdatedAt:   mainData.UpdatedAt,
	}
}

func ModelToMainMapping(dataModel []Rubbish) []entity.Main {
	var mainList []entity.Main
	for _, value := range dataModel {
		mainList = append(mainList, MapModelToMain(value))
	}
	return mainList
}
