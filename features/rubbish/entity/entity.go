package entity

import (
	"mime/multipart"
	"time"
)

type Main struct {
	Id          string
	Name        string `validate:"required"`
	TypeRubbish string `validate:"required"`
	PointPerKg  int    `validate:"required"`
	Description string
	ImageURL    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

type RubbishDataInterface interface {
	Create(data Main, image *multipart.FileHeader) error
	GetById(id string) (Main, error)
	GetByType(typeRubbish string) ([]Main, error)
	UpdateById(id string, updated Main, image *multipart.FileHeader) (data Main, err error)
	DeleteById(id string) error
	FindAllRubbish() ([]Main, error)
}

type UseCaseInterface interface {
	Create(data Main, image *multipart.FileHeader) error
	GetById(id string) (Main, error)
	GetByType(typeRubbish string) ([]Main, error)
	UpdateById(id string, updated Main, image *multipart.FileHeader) (data Main, err error)
	DeleteById(id string) error
	FindAllRubbish() ([]Main, error)
}
