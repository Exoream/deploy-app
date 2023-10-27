package controller

import (
	"recycle/features/detail_pickup/entity"
	"time"
)

type DetailPickupRespon struct {
	Id          string    `json:"id"`
	PickupId    string    `json:"pickup_id"`
	RubbishId   string    `json:"rubbish_id"`
	ItemWeight  float64       `json:"item_weight"`
	TotalPoints float64       `json:"total_points"`
	CreatedAt   time.Time `json:"created_at"`
}

func MainResponses(detail entity.Main) DetailPickupRespon {
	return DetailPickupRespon{
		Id:          detail.Id,
		PickupId:    detail.PickupId,
		RubbishId:   detail.RubbishId,
		ItemWeight:  detail.ItemWeight,
		TotalPoints: detail.TotalPoints,
		CreatedAt:   detail.CreatedAt,
	}
}

func MapModelsToController(detailPickups []entity.Main) []DetailPickupRespon {
	PickupGetAllData := make([]DetailPickupRespon, 0)
	for _, value := range detailPickups {
		pickupResponse := MainResponses(value)
		PickupGetAllData = append(PickupGetAllData, pickupResponse)
	}
	return PickupGetAllData
}