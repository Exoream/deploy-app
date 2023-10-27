package controller

import "recycle/features/pickup/entity"

type PickupRequest struct {
	Id         string `json:"id" form:"id"`
	Address    string `json:"address" form:"address"`
	Longitude  string `json:"longitude" form:"longitude"`
	Latitude   string `json:"latitude" form:"latitude"`
	PickupDate string `json:"pickup_date" form:"pickup_date"`
	UserId     string `json:"user_id" form:"user_id"`
	LocationId string `json:"location_id" form:"location_id"`
	ImageURL   string `json:"image_url" form:"image_url"`
}

func RequestMain(dataRequest PickupRequest) entity.Main {
	return entity.Main{
		Id:         dataRequest.Id,
		Address:    dataRequest.Address,
		Longitude:  dataRequest.Longitude,
		Latitude:   dataRequest.Latitude,
		PickupDate: dataRequest.PickupDate,
		LocationId: dataRequest.LocationId,
		ImageURL:   dataRequest.ImageURL,
	}
}
