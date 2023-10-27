package controller

import "recycle/features/rubbish/entity"

type RubbishRequest struct {
	Id          string `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	TypeRubbish string `json:"type_rubbish" form:"type_rubbish"`
	PointPerKg  int    `json:"point_per_kg" form:"point_per_kg"`
	Description string `json:"description" form:"description"`
	ImageURL    string `json:"image_url" form:"image_url"`
}

func RequestMain(dataRequest RubbishRequest) entity.Main {
	return entity.Main{
		Id:          dataRequest.Id,
		Name:        dataRequest.Name,
		TypeRubbish: dataRequest.TypeRubbish,
		PointPerKg:  dataRequest.PointPerKg,
		Description: dataRequest.Description,
		ImageURL:    dataRequest.ImageURL,
	}
}
