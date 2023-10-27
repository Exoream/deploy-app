package entity

import "time"

type Main struct {
	Id         string
	City       string `validate:"required"`
	Subdistric string `validate:"required"`
	PostalCode string `validate:"required"`
	Longitude  string `validate:"required"`
	Latitude   string `validate:"required"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

type LocationDataInterface interface {
	Create(data Main) error
	GetById(id string) (Main, error)
	GetByCity(city string) ([]Main, error)
	UpdateById(id string, updated Main) (data Main, err error)
	DeleteById(id string) error
	FindAllLocation() ([]Main, error)
}

type UseCaseInterface interface {
	Create(data Main) error
	GetById(id string) (Main, error)
	GetByCity(city string) ([]Main, error)
	UpdateById(id string, updated Main) (data Main, err error)
	DeleteById(id string) error
	FindAllLocation() ([]Main, error)
}
