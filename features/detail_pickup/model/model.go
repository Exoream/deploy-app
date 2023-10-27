package model

import (
	"recycle/features/rubbish/model"
	"time"

	"gorm.io/gorm"
)

type DetailPickup struct {
	Id          string         `gorm:"type:varchar(100);primaryKey;not null" json:"id"`
	PickupId    string         `gorm:"type:varchar(100);foreignKey:PickupId" json:"pickup_id"`
	RubbishId   string         `gorm:"type:varchar(100);foreignKey:RubbishId" json:"rubbish_id"`
	Rubbish     model.Rubbish  `gorm:"foreignKey:RubbishId"`
	ItemWeight  float64        `gorm:"not null" json:"item_weight"`
	TotalPoints float64        `gorm:"not null" json:"total_points"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
