package repository

import (
	"errors"
	"mime/multipart"
	"recycle/features/rubbish/entity"
	"recycle/features/rubbish/model"
	"recycle/storage"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type rubbishRepository struct {
	db *gorm.DB
}

func NewRubbishRepository(db *gorm.DB) entity.RubbishDataInterface {
	return &rubbishRepository{
		db: db,
	}
}

// Create implements rubbish.RubbishDataInterface.
func (u *rubbishRepository) Create(data entity.Main, image *multipart.FileHeader) error {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	dataInput := model.MapMainToModel(data)
	dataInput.Id = newUUID.String()

	imageURL, err := storage.UploadImageForRubbish(image)
	if err != nil {
		return err
	}

	dataInput.ImageURL = imageURL

	tx := u.db.Create(&dataInput)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// DeleteById implements rubbish.RubbishDataInterface.
func (u *rubbishRepository) DeleteById(id string) error {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	var rubbish model.Rubbish
	result := u.db.Where("id = ?", uuidID).Delete(&rubbish)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("failed to delete data, record not found")
	}

	return nil
}

// FindAllrubbish implements rubbish.RubbishDataInterface.
func (u *rubbishRepository) FindAllRubbish() ([]entity.Main, error) {
	var trash []model.Rubbish

	err := u.db.Find(&trash).Error
	if err != nil {
		return nil, err
	}

	var allrubbish []entity.Main = model.ModelToMainMapping(trash)

	return allrubbish, nil
}

// GetById implements rubbish.RubbishDataInterface.
func (u *rubbishRepository) GetById(id string) (entity.Main, error) {
	var rubbishData model.Rubbish
	result := u.db.Where("id = ?", id).First(&rubbishData)
	if result.Error != nil {
		return entity.Main{}, result.Error
	}

	var rubbishById = model.MapModelToMain(rubbishData)
	return rubbishById, nil
}

// UpdateById implements rubbish.RubbishDataInterface.
func (u *rubbishRepository) UpdateById(id string, updated entity.Main, image *multipart.FileHeader) (data entity.Main, err error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return entity.Main{}, err
	}

	var rubbishData model.Rubbish
	resultFind := u.db.First(&rubbishData, uuidID)
	if resultFind.Error != nil {
		return entity.Main{}, resultFind.Error
	}

	imageURL, uploadErr := storage.UploadImageForRubbish(image)
	if uploadErr != nil {
		return entity.Main{}, uploadErr
	}
	updated.ImageURL = imageURL

	u.db.Model(&rubbishData).Updates(model.MapMainToModel(updated))

	data = model.MapModelToMain(rubbishData)
	return data, nil
}

// GetByType implements entity.RubbishDataInterface.
func (u *rubbishRepository) GetByType(typeRubbish string) ([]entity.Main, error) {
	var typeData []model.Rubbish
	result := u.db.Where("type_rubbish = ?", typeRubbish).Find(&typeData)
	if result.Error != nil {
		return nil, result.Error
	}

	var rubbish []entity.Main
	rubbish = model.ModelToMainMapping(typeData)
	return rubbish, nil
}
