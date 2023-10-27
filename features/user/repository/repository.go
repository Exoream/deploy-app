package repository

import (
	"errors"
	"recycle/email"
	pick "recycle/features/pickup/model"
	"recycle/features/user/entity"
	"recycle/features/user/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) entity.UserDataInterface {
	return &userRepository{
		db: db,
	}
}

// Create implements user.UserDataInterface.
func (u *userRepository) Create(user entity.Main) (string, error) {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	dataInput := model.MapMainToModel(user)
	dataInput.Id = newUUID.String()
	uniqueToken := email.GenerateUniqueToken()
	dataInput.VerificationToken = uniqueToken

	tx := u.db.Create(&dataInput)
	if tx.Error != nil {
		return "", tx.Error
	}

	return uniqueToken, nil
}

// CheckLogin implements user.UserDataInterface.
func (u *userRepository) CheckLogin(email string, password string) (entity.Main, error) {
	var data model.User

	tx := u.db.Where("email = ?", email).First(&data)
	if tx.Error != nil {
		return entity.Main{}, tx.Error
	}
	dataMain := model.MapModelToMain(data)
	return dataMain, nil
}

// GetById implements user.UserDataInterface.
func (u *userRepository) GetById(id string) (entity.Main, error) {
	var userData model.User

	// Gunakan Preload untuk memuat data pickup terkait.
	result := u.db.Preload("Pickups").Where("id = ?", id).First(&userData)
	if result.Error != nil {
		return entity.Main{}, result.Error
	}

	var userById = model.MapModelToMain(userData)
	userById.Pickups = pick.ModelToMainMapping(userData.Pickups)
	return userById, nil
}

// UpdateById implements user.UserDataInterface.
func (u *userRepository) UpdateById(id string, updated entity.Main) (data entity.Main, err error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return entity.Main{}, err
	}

	var usersData model.User
	resultFind := u.db.First(&usersData, uuidID)
	if resultFind.Error != nil {
		return entity.Main{}, resultFind.Error
	}

	u.db.Model(&usersData).Updates(model.MapMainToModel(updated))

	data = model.MapModelToMain(usersData)
	return data, nil
}

// Delete implements user.UserDataInterface.
func (u *userRepository) DeleteById(id string) error {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	var user model.User
	result := u.db.Where("id = ?", uuidID).Delete(&user)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("failed to delete data, record not found")
	}

	return nil
}

// FindAllUsers implements user.UserDataInterface.
func (u *userRepository) FindAllUsers() ([]entity.Main, error) {
	var users []model.User

	err := u.db.Preload("Pickups").Find(&users).Error
	if err != nil {
		return nil, err
	}

	var allUser []entity.Main = model.ModelToMainMapping(users)

	return allUser, nil
}

// GetByVerificationToken implements entity.UserDataInterface.
func (u *userRepository) GetByVerificationToken(token string) (entity.Main, error) {
	var userData model.User
	result := u.db.Where("verification_token = ?", token).First(&userData)
	if result.Error != nil {
		return entity.Main{}, result.Error
	}

	var userToken = model.MapModelToMain(userData)
	return userToken, nil
}

// UpdateIsVerified implements entity.UserDataInterface.
func (u *userRepository) UpdateIsVerified(userID string, isVerified bool) error {
	uuidID, err := uuid.Parse(userID)
	if err != nil {
		return err
	}

	var user model.User
	result := u.db.First(&user, uuidID)
	if result.Error != nil {
		return result.Error
	}

	user.IsVerified = isVerified
	result = u.db.Save(&user)

	return result.Error
}

// GetEmailByID implements entity.UserDataInterface.
func (u *userRepository) GetEmailByID(userID string) (string, error) {
	var userEmail string
	var user model.User
	if err := u.db.Where("id = ?", userID).First(&user).Error; err != nil {
        return "", err
    }

	userEmail = user.Email
    return userEmail, nil
}
