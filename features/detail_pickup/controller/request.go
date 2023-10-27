package controller

import "recycle/features/detail_pickup/entity"

type DetailPickupDetail struct {
	RubbishID  string  `json:"rubbish_id"`
	ItemWeight float64 `json:"item_weight"`
}

type DetailPickupRequest struct {
	PickupID      string               `json:"pickup_id"`
	PickupDetails []DetailPickupDetail `json:"pickup_details"`
}

func RequestDetailPickup(dataRequest DetailPickupRequest) []entity.Main {
	var detailPickups []entity.Main

	for _, detail := range dataRequest.PickupDetails {
		detailPickup := entity.Main{
			PickupId:    dataRequest.PickupID,
			RubbishId:   detail.RubbishID,
			ItemWeight:  detail.ItemWeight,
			TotalPoints: 0,
		}
		detailPickups = append(detailPickups, detailPickup)
	}
	return detailPickups
}
