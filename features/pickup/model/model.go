package model

import (
	pickup "recycle/features/detail_pickup/model"
	"recycle/features/location/model"
	"time"

	"gorm.io/gorm"
)

type Pickup struct {
	Id           string                `gorm:"type:varchar(100);primaryKey;not null" json:"id" form:"id"`
	Address      string                `gorm:"type:longtext;not null" json:"address" form:"address"`
	Longitude    string                `gorm:"type:varchar(255);not null" json:"longitude" form:"longitude"`
	Latitude     string                `gorm:"type:varchar(255);not null" json:"latitude" form:"latitude"`
	PickupDate   string                `gorm:"type:date;not null" json:"pickup_date" form:"pickup_date"`
	Status       string                `gorm:"type:varchar(20);not null" json:"status" form:"status"`
	UserId       string                `gorm:"type:varchar(100);foreignKey:UserId" json:"user_id" form:"user_id"`
	LocationId   string                `gorm:"type:varchar(100);foreignKey:LocationId" json:"location_id" form:"location_id"`
	ImageURL     string                `gorm:"type:varchar(255)" json:"image_url" form:"image_url"`
	Location     model.Location        `gorm:"foreignKey:LocationId"`
	CreatedAt    time.Time             `json:"created_at"`
	UpdatedAt    time.Time             `json:"updated_at"`
	DeletedAt    gorm.DeletedAt        `gorm:"index"`
	DetailPickup []pickup.DetailPickup `gorm:"foreignKey:PickupId"`
}
