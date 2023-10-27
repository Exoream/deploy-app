package model

import (
	"time"

	"gorm.io/gorm"
)

type Location struct {
	Id         string         `gorm:"type:varchar(100);primaryKey;not null" json:"id"`
	City       string         `gorm:"type:varchar(255);not null" json:"city"`
	Subdistric string         `gorm:"type:varchar(255);not null;unique" json:"subdistric"`
	PostalCode string         `gorm:"type:varchar(10)not null" json:"postal_code"`
	Longitude  string         `gorm:"type:varchar(255);not null" json:"longitude"`
	Latitude   string         `gorm:"type:varchar(255);not null" json:"latitude"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
