package repository

import (
	"recycle/features/detail_pickup/entity"
	"recycle/features/detail_pickup/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type detailPickupRepository struct {
	db *gorm.DB
}

func NewDetailPickupRepository(db *gorm.DB) entity.DetailPickupDataInterface {
	return &detailPickupRepository{
		db: db,
	}
}

// Create implements entity.DetailPickupDataInterface.
func (u *detailPickupRepository) Create(data []entity.Main) error {
	tx := u.db.Begin()

    for _, detail := range data {
        newUUID, err := uuid.NewRandom()
        if err != nil {
            tx.Rollback()
            return err
        }

        detailInput := model.MapMainToModel(detail)
        detailInput.Id = newUUID.String()

        if err := tx.Create(&detailInput).Error; err != nil {
            tx.Rollback()
            return err
        }
    }

    tx.Commit()
    return nil
}

// FindAllDetailPickup implements entity.DetailPickupDataInterface.
func (u *detailPickupRepository) FindAllDetailPickup() ([]entity.Main, error) {
	var pickup []model.DetailPickup

	err := u.db.Find(&pickup).Error
	if err != nil {
		return nil, err
	}

	var allPickup []entity.Main
	allPickup = model.ModelToMainMapping(pickup)

	return allPickup, nil
}

// Implementasi antarmuka DetailPickupDataInterface dalam repositori Anda
func (u *detailPickupRepository) BeginTransaction() (*gorm.DB, error) {
    // Membuka transaksi dengan GORM
    tx := u.db.Begin()
    if tx.Error != nil {
        return nil, tx.Error
    }
    return tx, nil
}

