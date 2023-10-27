package model

import (
	"recycle/features/pickup/model"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id                string         `gorm:"type:varchar(100);primaryKey;not null" json:"id"`
	Name              string         `gorm:"type:varchar(255);not null" json:"name"`
	Email             string         `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password          string         `gorm:"type:text;not null" json:"password"`
	Gender            string         `gorm:"type:enum('male','female')" json:"gender"`
	Age               int            `json:"age"`
	Address           string         `gorm:"type:longtext" json:"address"`
	SaldoPoints       int            `json:"saldo_points"`
	Role              string         `gorm:"type:varchar(10)" json:"role"`
	IsVerified        bool           `gorm:"default:false" json:"is_verified"`
	VerificationToken string         `gorm:"type:varchar(255)" json:"verification_token"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	Pickups           []model.Pickup `gorm:"foreignKey:UserId"`
}
