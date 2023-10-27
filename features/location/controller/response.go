package controller

import "recycle/features/location/entity"

type LocationRespon struct {
	Id         string `json:"id"`
	City       string `json:"city"`
	Subdistric string `json:"subdistric"`
	PostalCode string `json:"postal_code"`
	Longitude  string `json:"longitude"`
	Latitude   string `json:"latitude"`
}

func MainResponse(dataMain entity.Main) LocationRespon {
	return LocationRespon{
		Id:         dataMain.Id,
		City:       dataMain.City,
		Subdistric: dataMain.Subdistric,
		PostalCode: dataMain.PostalCode,
		Longitude:  dataMain.Longitude,
		Latitude:   dataMain.Latitude,
	}
}
