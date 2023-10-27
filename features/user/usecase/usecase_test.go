package usecase

import (
	"errors"
	"testing"

	"recycle/features/user/entity"
	"recycle/features/user/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateSuccess(t *testing.T) {
	mockRepo := new(mocks.UserDataInterface)
	userUC := NewUserUsecase(mockRepo)

	testPassword := "securePassword"
	testData := entity.Main{
		Name:     "John Doe",
		Email:    "johndoe@gmail.com",
		Password: testPassword, // Use the unhashed password
		Gender:   "male",
		Age:      20,
		Address:  "madagaskar",
		Role:     "user",
	}

	mockRepo.On("Create", mock.Anything).Return("", nil)

	uniqueToken, err := userUC.Create(testData)

	assert.Nil(t, err)
	assert.Equal(t, "", uniqueToken)
	mockRepo.AssertExpectations(t)
}

func TestCreateError(t *testing.T) {
	mockRepo := new(mocks.UserDataInterface)
	userUC := NewUserUsecase(mockRepo)

	// Prepare test data
	testData := entity.Main{
		Name:     "John Doe",
		Email:    "johndoe@gmail.com",
		Password: "securePassword",
		Gender:   "male",
		Age:      20,
		Address:  "madagaskar",
		Role:     "user",
	}

	mockRepo.On("Create", mock.Anything).Return("", errors.New("some error message")).Once()
	token, err := userUC.Create(testData)

	assert.NotNil(t, err)
	assert.Equal(t, "some error message", err.Error())
	assert.Equal(t, "", token)
	mockRepo.AssertExpectations(t)
}

func TestCreateErrorValidation(t *testing.T) {
	mockRepo := new(mocks.UserDataInterface)
	userUC := NewUserUsecase(mockRepo)

	testData := entity.Main{
		Name:     "",
		Email:    "",
		Password: "",
		Gender:   "",
	}

	token, err := userUC.Create(testData)
	assert.NotNil(t, err)
	assert.Empty(t, token)
}

func TestCheckLoginErrorValidation(t *testing.T) {
	mockRepo := new(mocks.UserDataInterface)
	userUC := NewUserUsecase(mockRepo)

	email := ""
	password := ""

	userData, token, err := userUC.CheckLogin(email, password)

	assert.NotNil(t, err)
	assert.Empty(t, userData)
	assert.Empty(t, token)

	mockRepo.AssertExpectations(t)
}

func TestCheckLoginErrorCheckPasswordHash(t *testing.T) {
	mockRepo := new(mocks.UserDataInterface)
	userUC := NewUserUsecase(mockRepo)

	// Prepare test data
	email := "user@example.com"
	password := "password123"

	mockUserData := entity.Main{
		Email:    email,
		Password: password,
	}

	mockRepo.On("CheckLogin", email, password).Return(mockUserData, nil)

	userData, token, err := userUC.CheckLogin(email, password)

	assert.NotNil(t, err)
	assert.Empty(t, userData)
	assert.Empty(t, token)

	mockRepo.AssertExpectations(t)
}

func TestGetByIdSuccess(t *testing.T) {
	mockRepo := new(mocks.UserDataInterface)
	userUC := NewUserUsecase(mockRepo)

	userID := "some_user_id"

	mockUserData := entity.Main{
		Id: userID,
	}

	mockRepo.On("GetById", userID).Return(mockUserData, nil)
	data, _ := userUC.GetById(userID)

	assert.Equal(t, mockUserData, data)
	mockRepo.AssertExpectations(t)
}

func TestGetByIdErrorInvalidID(t *testing.T) {
	mockRepo := new(mocks.UserDataInterface)
	userUC := NewUserUsecase(mockRepo)

	userID := ""

	data, err := userUC.GetById(userID)
	assert.NotNil(t, err)
	assert.Empty(t, data)
}

func TestUpdateByIdSuccess(t *testing.T) {
	mockRepo := new(mocks.UserDataInterface)
	userUC := NewUserUsecase(mockRepo)

	userID := "some_user_id"
	updatedData := entity.Main{
		Id: userID,
	}

	mockUserData := entity.Main{
		Id: "updated",
	}

	mockRepo.On("UpdateById", userID, updatedData).Return(mockUserData, nil)

	data, err := userUC.UpdateById(userID, updatedData)

	assert.Nil(t, err)
	assert.Equal(t, mockUserData, data)

	mockRepo.AssertExpectations(t)
}

func TestUpdateByIdErrorInvalidID(t *testing.T) {
	mockRepo := new(mocks.UserDataInterface)
	userUC := NewUserUsecase(mockRepo)

	userID := ""
	updatedData := entity.Main{
		Id: userID,
	}

	data, err := userUC.UpdateById(userID, updatedData)
	assert.NotNil(t, err)
	assert.Empty(t, data)
}

func TestUpdateByIdErrorHashPassword(t *testing.T) {
	mockRepo := new(mocks.UserDataInterface)
	userUC := NewUserUsecase(mockRepo)

	userID := "some_user_id"
	updatedData := entity.Main{
		Password: "invalid_password",
	}

	mockRepo.On("UpdateById", userID, updatedData).Return(entity.Main{}, errors.New("mock error"))
	data, err := userUC.UpdateById(userID, updatedData)

	assert.NotNil(t, err)
	assert.Empty(t, data)
}

func TestUpdateByIdError(t *testing.T) {
	mockRepo := new(mocks.UserDataInterface)
	userUC := NewUserUsecase(mockRepo)

	// Prepare test data
	userID := "some_user_id"
	updatedData := entity.Main{
		Id: "",
	}

	mockRepo.On("UpdateById", userID, updatedData).Return(entity.Main{}, errors.New("mock error"))
	data, err := userUC.UpdateById(userID, updatedData)

	assert.NotNil(t, err)
	assert.Empty(t, data)
}

func TestDeleteByIdSuccess(t *testing.T) {
	mockRepo := new(mocks.UserDataInterface)
	userUC := NewUserUsecase(mockRepo)

	userID := "some_user_id"
	mockRepo.On("DeleteById", userID).Return(nil)

	err := userUC.DeleteById(userID)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteByIdErrorInvalidID(t *testing.T) {
	mockRepo := new(mocks.UserDataInterface)
	userUC := NewUserUsecase(mockRepo)

	userID := ""

	err := userUC.DeleteById(userID)
	assert.NotNil(t, err)
}

func TestDeleteByIdError(t *testing.T) {
	mockRepo := new(mocks.UserDataInterface)
	userUC := NewUserUsecase(mockRepo)

	userID := "some_user_id"
	mockRepo.On("DeleteById", userID).Return(errors.New("mock error"))

	err := userUC.DeleteById(userID)
	assert.NotNil(t, err)
}

func TestFindAllUsersSuccess(t *testing.T) {
	mockRepo := new(mocks.UserDataInterface)
	userUC := NewUserUsecase(mockRepo)

	// Mocked user data list
	mockUserList := []entity.Main{
		{
			Id:    "satu",
			Name:  "User 1",
			Email: "user1@example.com",
			// Isi dengan data lain yang sesuai
		},
		{
			Id:    "dua",
			Name:  "User 2",
			Email: "user2@example.com",
		},
	}

	mockRepo.On("FindAllUsers").Return(mockUserList, nil)
	users, err := userUC.FindAllUsers()

	assert.Nil(t, err)
	assert.Equal(t, mockUserList, users)

	mockRepo.AssertExpectations(t)
}

func TestFindAllUsersError(t *testing.T) {
	mockRepo := new(mocks.UserDataInterface)
	userUC := NewUserUsecase(mockRepo)

	mockRepo.On("FindAllUsers").Return(nil, errors.New("mock error"))
	users, err := userUC.FindAllUsers()

	assert.NotNil(t, err)
	assert.Nil(t, users)
}

func TestGetByVerificationTokenSuccess(t *testing.T) {
	mockRepo := new(mocks.UserDataInterface)
	userUC := NewUserUsecase(mockRepo)

	token := "some_verification_token"

	mockUserData := entity.Main{
		Id:                "1",
		Name:              "John Doe",
		Email:             "johndoe@example.com",
		VerificationToken: token,
	}

	mockRepo.On("GetByVerificationToken", token).Return(mockUserData, nil)
	data, err := userUC.GetByVerificationToken(token)

	assert.Nil(t, err)
	assert.Equal(t, mockUserData, data)

	mockRepo.AssertExpectations(t)
}

func TestGetByVerificationInvalidToken(t *testing.T) {
	mockRepo := new(mocks.UserDataInterface)
	userUC := NewUserUsecase(mockRepo)

	token := ""

	data, err := userUC.GetByVerificationToken(token)
	assert.NotNil(t, err)
	assert.Empty(t, data)
}

func TestGetByVerificationTokenError(t *testing.T) {
	mockRepo := new(mocks.UserDataInterface)
	userUC := NewUserUsecase(mockRepo)

	token := "some_verification_token"

	mockRepo.On("GetByVerificationToken", token).Return(entity.Main{}, errors.New("mock error"))

	data, err := userUC.GetByVerificationToken(token)

	assert.NotNil(t, err)
	assert.Empty(t, data)
}

func TestUpdateIsVerifiedSuccess(t *testing.T) {
	mockRepo := new(mocks.UserDataInterface)
	userUC := NewUserUsecase(mockRepo)

	userID := "some_user_id"
	isVerified := true

	mockRepo.On("UpdateIsVerified", userID, isVerified).Return(nil)

	err := userUC.UpdateIsVerified(userID, isVerified)
	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}

func TestUpdateIsVerifiedInvalidUserID(t *testing.T) {
	mockRepo := new(mocks.UserDataInterface)
	userUC := NewUserUsecase(mockRepo)

	userID := ""

	err := userUC.UpdateIsVerified(userID, true)
	assert.NotNil(t, err)
}

func TestGetEmailByIDSuccess(t *testing.T) {
	mockRepo := new(mocks.UserDataInterface)
	userUC := NewUserUsecase(mockRepo)

	userID := "some_user_id"
	mockEmail := "user@example.com"

	mockRepo.On("GetEmailByID", userID).Return(mockEmail, nil)

	email, err := userUC.GetEmailByID(userID)

	assert.Nil(t, err)
	assert.Equal(t, mockEmail, email)

	mockRepo.AssertExpectations(t)
}

func TestCreatePasswordTooShort(t *testing.T) {
	mockRepo := new(mocks.UserDataInterface)
	userUC := NewUserUsecase(mockRepo)

	testData := entity.Main{
		Name:     "John Doe",
		Email:    "johndoe@gmail.com",
		Password: "short",
		Gender:   "male",
		Age:      20,
		Address:  "madagaskar",
		Role:     "user",
	}

	uniqueToken, err := userUC.Create(testData)

	expectedErrorMessage := "your password too short"
	assert.NotNil(t, err)
	assert.Empty(t, uniqueToken)
	assert.Equal(t, expectedErrorMessage, err.Error())

	mockRepo.AssertExpectations(t)
}

func TestCheckLoginFailedLogin(t *testing.T) {
	mockRepo := new(mocks.UserDataInterface)
	userUC := NewUserUsecase(mockRepo)

	validEmail := "user@example.com"
	validPassword := "password123"

	mockRepo.On("CheckLogin", validEmail, validPassword).Return(entity.Main{}, errors.New("check login failed"))

	_, _, err := userUC.CheckLogin(validEmail, validPassword)

	assert.NotNil(t, err)
	assert.Equal(t, "check login failed", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestUserUseCase_GetEmailByID_Error(t *testing.T) {
    mockRepo := new(mocks.UserDataInterface)
    userUC := NewUserUsecase(mockRepo)

    userID := "user123"

    mockRepo.On("GetEmailByID", userID).Return("", errors.New("user not found"))

    _, err := userUC.GetEmailByID(userID)

    assert.NotNil(t, err)
    assert.Equal(t, "user not found", err.Error())

    mockRepo.AssertExpectations(t)
}