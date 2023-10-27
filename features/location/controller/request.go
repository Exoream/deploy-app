package controller

import "recycle/features/location/entity"

type LocationRequest struct {
	Id         string `json:"id"`
	City       string `json:"city"`
	Subdistric string `json:"subdistric"`
	PostalCode string `json:"postal_code"`
	Longitude  string `json:"longitude"`
	Latitude   string `json:"latitude"`
}

func RequestMain(dataRequest LocationRequest) entity.Main {
	return entity.Main{
		Id:         dataRequest.Id,
		City:       dataRequest.City,
		Subdistric: dataRequest.Subdistric,
		PostalCode: dataRequest.PostalCode,
		Longitude:  dataRequest.Longitude,
		Latitude:   dataRequest.Latitude,
	}
}
