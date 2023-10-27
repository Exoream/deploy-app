package entity

import (
	"time"

	"gorm.io/gorm"
)

type Main struct {
	Id          string
	PickupId    string  `validate:"required"`
	RubbishId   string  `validate:"required"`
	ItemWeight  float64 `validate:"required"`
	TotalPoints float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

type DetailPickupDataInterface interface {
	Create(data []Main) error
	FindAllDetailPickup() ([]Main, error)
	BeginTransaction() (*gorm.DB, error)
}

type UseCaseInterface interface {
	Create(data []Main) (int, error)
	FindAllDetailPickup() ([]Main, error)
}
