package entity

import (
	"recycle/features/pickup/entity"
	"time"
)

type Main struct {
	Id                string
	Name              string `validate:"required"`
	Email             string `validate:"required,email"`
	Password          string `validate:"required"`
	Gender            string `validate:"oneof=male female"`
	Age               int
	Address           string
	SaldoPoints       int
	Role              string
	IsVerified        bool
	VerificationToken string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time
	Pickups           []entity.Main
}

type UserDataInterface interface {
	Create(data Main) (string, error)
	CheckLogin(email, password string) (Main, error)
	GetById(id string) (Main, error)
	UpdateById(id string, updated Main) (data Main, err error)
	DeleteById(id string) error
	FindAllUsers() ([]Main, error)
	GetByVerificationToken(token string) (Main, error)
	UpdateIsVerified(userID string, isVerified bool) error
	GetEmailByID(userID string) (string, error)
}

type UseCaseInterface interface {
	Create(data Main) (string, error)
	CheckLogin(email, password string) (Main, string, error)
	GetById(id string) (Main, error)
	UpdateById(id string, updated Main) (data Main, err error)
	DeleteById(id string) error
	FindAllUsers() ([]Main, error)
	GetByVerificationToken(token string) (Main, error)
	UpdateIsVerified(userID string, isVerified bool) error
	GetEmailByID(userID string) (string, error)
}
