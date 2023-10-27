package entity

import (
	"mime/multipart"
	"time"
)

type Main struct {
	Id         string
	Address    string `validate:"required"`
	Longitude  string `validate:"required"`
	Latitude   string `validate:"required"`
	PickupDate string `validate:"required"`
	Status     string
	UserId     string
	LocationId string `validate:"required"`
	ImageURL   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

type PickupDataInterface interface {
	Create(data Main, image *multipart.FileHeader) error
	UpdateById(id string, updated Main, image *multipart.FileHeader) (data Main, err error)
	DeleteById(id string) error
	GetById(id string) (Main, error)
	FindAllPickup() ([]Main, error)
	GetByStatus(status string) ([]Main, error)
	UpdateStatus(pickupID, newStatus string) error
}

type UseCaseInterface interface {
	Create(data Main, image *multipart.FileHeader) error
	UpdateById(id string, updated Main, image *multipart.FileHeader) (data Main, err error)
	DeleteById(id string) error
	GetById(id string) (Main, error)
	FindAllPickup() ([]Main, error)
	GetByStatus(status string) ([]Main, error)
}
