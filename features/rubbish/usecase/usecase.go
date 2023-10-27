package usecase

import (
	"errors"
	"mime/multipart"
	"recycle/features/rubbish/entity"

	"github.com/go-playground/validator/v10"
)

type rubbishUseCase struct {
	rubbishRepo entity.RubbishDataInterface
	validate    *validator.Validate
}

func NewRubbishUsecase(rubbishRepo entity.RubbishDataInterface) entity.UseCaseInterface {
	return &rubbishUseCase{
		rubbishRepo: rubbishRepo,
		validate:    validator.New(),
	}
}

// Create implements rubbish.UseCaseInterface.
func (uc *rubbishUseCase) Create(data entity.Main, image *multipart.FileHeader) error {
	errValidate := uc.validate.Struct(data)
	if errValidate != nil {
		return errValidate
	}

	if image.Size > 10*1024*1024 {
		return errors.New("image file size should be less than 10 MB")
	}

	err := uc.rubbishRepo.Create(data, image)
	if err != nil {
		return err
	}

	return nil
}

// DeleteById implements rubbish.UseCaseInterface.
func (uc *rubbishUseCase) DeleteById(id string) error {
	if id == "" {
		return errors.New("id not found")
	}

	err := uc.rubbishRepo.DeleteById(id)
	if err != nil {
		return err
	}

	return nil
}

// FindAllrubbish implements rubbish.UseCaseInterface.
func (uc *rubbishUseCase) FindAllRubbish() ([]entity.Main, error) {
	rubbish, err := uc.rubbishRepo.FindAllRubbish()
	if err != nil {
		return nil, err
	}
	return rubbish, nil
}

// GetById implements rubbish.UseCaseInterface.
func (uc *rubbishUseCase) GetById(id string) (entity.Main, error) {
	if id == "" {
		return entity.Main{}, errors.New("invalid id")
	}

	idRubbish, err := uc.rubbishRepo.GetById(id)
	return idRubbish, err
}

// UpdateById implements rubbish.UseCaseInterface.
func (uc *rubbishUseCase) UpdateById(id string, updated entity.Main, image *multipart.FileHeader) (data entity.Main, err error) {
	if id == "" {
		return entity.Main{}, errors.New("id not found")
	}

	if image != nil && image.Size > 10*1024*1024 {
		return entity.Main{}, errors.New("image file size should be less than 10 MB")
	}

	data, err = uc.rubbishRepo.UpdateById(id, updated, image)
	if err != nil {
		return entity.Main{}, err
	}
	return data, nil
}

// GetByType implements entity.UseCaseInterface.
func (uc *rubbishUseCase) GetByType(typeRubbish string) ([]entity.Main, error) {
	if typeRubbish == "" {
		return nil, errors.New("typeRubbish parameter is required")
	}

	typeName, err := uc.rubbishRepo.GetByType(typeRubbish)
	return typeName, err
}
