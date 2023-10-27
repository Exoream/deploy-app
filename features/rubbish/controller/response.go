package controller

import (
	"recycle/features/rubbish/entity"
)

type RubbishRespon struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	TypeRubbish string `json:"type_rubbish"`
	PointPerKg  int    `json:"point_per_kg"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

func MainResponse(dataMain entity.Main) RubbishRespon {
	return RubbishRespon{
		Id:          dataMain.Id,
		Name:        dataMain.Name,
		TypeRubbish: dataMain.TypeRubbish,
		PointPerKg:  dataMain.PointPerKg,
		Description: dataMain.Description,
		ImageURL:    dataMain.ImageURL,
	}
}
