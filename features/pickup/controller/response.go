package controller

import "recycle/features/pickup/entity"

type PickupRespon struct {
	Id         string `json:"id"`
	Address    string `json:"address"`
	Longitude  string `json:"longitude"`
	Latitude   string `json:"latitude"`
	PickupDate string `json:"pickup_date"`
	LocationId string `json:"location_id"`
	ImageURL   string `json:"image_url"`
	Status     string `json:"status"`
}

type PickupResponForGetAll struct {
	Id         string `json:"id"`
	Address    string `json:"address"`
	Longitude  string `json:"longitude"`
	Latitude   string `json:"latitude"`
	PickupDate string `json:"pickup_date"`
	UserId     string `json:"user_id"`
	LocationId string `json:"location_id"`
	ImageURL   string `json:"image_url"`
	Status     string `json:"status"`
}

func MainResponse(dataMain entity.Main) PickupRespon {
	return PickupRespon{
		Id:         dataMain.Id,
		Address:    dataMain.Address,
		Longitude:  dataMain.Longitude,
		Latitude:   dataMain.Latitude,
		PickupDate: dataMain.PickupDate,
		LocationId: dataMain.LocationId,
		ImageURL:   dataMain.ImageURL,
		Status:     dataMain.Status,
	}
}

func MapModelToController(pickups []entity.Main) []PickupRespon {
	pickupResponses := make([]PickupRespon, 0)
	for _, pickup := range pickups {
		pickupResponse := MainResponse(pickup)
		pickupResponses = append(pickupResponses, pickupResponse)
	}
	return pickupResponses
}

func MainResponses(dataMain entity.Main) PickupResponForGetAll {
	return PickupResponForGetAll{
		Id:         dataMain.Id,
		Address:    dataMain.Address,
		Longitude:  dataMain.Longitude,
		Latitude:   dataMain.Latitude,
		PickupDate: dataMain.PickupDate,
		UserId:     dataMain.UserId,
		LocationId: dataMain.LocationId,
		ImageURL:   dataMain.ImageURL,
		Status:     dataMain.Status,
	}
}

func MapModelsToController(pickups []entity.Main) []PickupResponForGetAll {
	PickupGetAllData := make([]PickupResponForGetAll, 0)
	for _, value := range pickups {
		pickupResponse := MainResponses(value)
		PickupGetAllData = append(PickupGetAllData, pickupResponse)
	}
	return PickupGetAllData
}
