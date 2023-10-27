package usecase

import (
	"errors"
	"recycle/features/location/entity"

	"github.com/go-playground/validator/v10"
)

type locationUseCase struct {
	locationRepo entity.LocationDataInterface
	validate     *validator.Validate
}

func NewLocationUsecase(locationRepo entity.LocationDataInterface) entity.UseCaseInterface {
	return &locationUseCase{
		locationRepo: locationRepo,
		validate:     validator.New(),
	}
}

// Create implements entity.UseCaseInterface.
func (uc *locationUseCase) Create(data entity.Main) error {
	errValidate := uc.validate.Struct(data)
	if errValidate != nil {
		return errValidate
	}

	err := uc.locationRepo.Create(data)
	if err != nil {
		return err
	}

	return nil
}

// DeleteById implements entity.UseCaseInterface.
func (uc *locationUseCase) DeleteById(id string) error {
	if id == "" {
		return errors.New("id not found")
	}

	err := uc.locationRepo.DeleteById(id)
	if err != nil {
		return err
	}

	return nil
}

// FindAllLocation implements entity.UseCaseInterface.
func (uc *locationUseCase) FindAllLocation() ([]entity.Main, error) {
	location, err := uc.locationRepo.FindAllLocation()
	if err != nil {
		return nil, err
	}
	return location, nil
}

// GetById implements entity.UseCaseInterface.
func (uc *locationUseCase) GetById(id string) (entity.Main, error) {
	if id == "" {
		return entity.Main{}, errors.New("invalid id")
	}

	idLocation, err := uc.locationRepo.GetById(id)
	return idLocation, err
}

// GetByCity implements entity.UseCaseInterface.
func (uc *locationUseCase) GetByCity(city string) ([]entity.Main, error) {
	if city == "" {
		return nil, errors.New("City parameter is required")
	}

	cityName, err := uc.locationRepo.GetByCity(city)
	return cityName, err
}

// UpdateById implements entity.UseCaseInterface.
func (uc *locationUseCase) UpdateById(id string, updated entity.Main) (data entity.Main, err error) {
	if id == "" {
		return entity.Main{}, errors.New("id not found")
	}

	data, err = uc.locationRepo.UpdateById(id, updated)
	if err != nil {
		return entity.Main{}, err
	}
	return data, nil
}
