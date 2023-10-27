package controller

import (
	"io/ioutil"
	"net/http"
	"recycle/app/middlewares"
	"recycle/email"
	user "recycle/features/user/entity"
	"recycle/helper"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUseCase user.UseCaseInterface
}

func NewUserControllers(uc user.UseCaseInterface) *UserController {
	return &UserController{
		userUseCase: uc,
	}
}

func (uco *UserController) CreateUser(c echo.Context) error {
	// Bind data
	dataInput := UserRequest{}
	errBind := c.Bind(&dataInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("error bind data"))
	}

	data := RequestMain(dataInput)

	uniqueToken, errCreate := uco.userUseCase.Create(data)
	if errCreate != nil {
		if strings.Contains(errCreate.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse(errCreate.Error()))
		} else {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("failed to create data "+errCreate.Error()))
		}
	}

	verificationLink := "http://localhost:8080/verify?token=" + uniqueToken
	emailTemplate, err := ioutil.ReadFile("email/templates/account_registration.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to read email template"))
	}

	_, errEmail := email.SendEmailSMTP([]string{data.Email}, string(emailTemplate), verificationLink)
	if errEmail != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to send verification email"))
	}

	return c.JSON(http.StatusCreated, helper.SuccesResponses("success create data"))
}

func (uco *UserController) Login(c echo.Context) error {
	// Bind data
	var login UserLogin
	errBind := c.Bind(&login)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("eror bind"))
	}

	// Memanggil func di usecase
	user, token, err := uco.userUseCase.CheckLogin(login.Email, login.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("login failed"))
	}

	if !user.IsVerified {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("account not verified"))
	}

	response := LoginResponse(user.Id, user.Email, token)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("login successful", response))
}

func (uco *UserController) GetUser(c echo.Context) error {
	// Extra token dari id
	idToken := middlewares.ExtractToken(c)
	if idToken == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
	}

	// Mendapatkan email pengguna dari token
	role, err := middlewares.ExtractRole(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("error extracting role"))
	}

	idParamStr := c.Param("id")
	idParam, errId := uuid.Parse(idParamStr)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid UUID format"))
	}

	// Periksa apakah ID dari token sama dengan ID dari parameter URL.
	if idToken.String() == idParam.String() || role == "admin" {
		// Jika sesuai, izinkan akses ke profil pengguna.
		result, err := uco.userUseCase.GetById(idParam.String())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error reading data"))
		}

		var usersResponse = MainResponse(result)

		return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get profile", usersResponse))
	}

	// Jika tidak ada izin yang sesuai, kembalikan pesan "unauthorized."
	return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
}

func (uco *UserController) Update(c echo.Context) error {
	// Extra token from id
	idToken := middlewares.ExtractToken(c)
	if idToken == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
	}

	// Extract the user's role from the token
	role, err := middlewares.ExtractRole(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("error extracting role"))
	}

	idParamStr := c.Param("id")
	idParam, errId := uuid.Parse(idParamStr)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid UUID format"))
	}

	if idToken.String() == idParam.String() || role == "admin" {
		userReq := UserRequest{}
		errBind := c.Bind(&userReq)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("error binding data"))
		}

		userMain := RequestMain(userReq)

		if role != "admin" && userMain.SaldoPoints != 0 {
			return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("you do not have permission to update saldo_points"))
		}

		if role != "admin" && userMain.Role != "" {
			return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("you do not have permission to update role"))
		}

		data, err := uco.userUseCase.UpdateById(idParam.String(), userMain)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success update data", MainResponse(data)))
	} else {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
	}
}

func (uco *UserController) Delete(c echo.Context) error {
	// Extra token dari id
	idToken := middlewares.ExtractToken(c)
	if idToken == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
	}

	// Mendapatkan email pengguna dari token
	role, err := middlewares.ExtractRole(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("error extracting role"))
	}

	idParamStr := c.Param("id")
	idParam, errId := uuid.Parse(idParamStr)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid UUID format"))
	}

	if idToken.String() == idParam.String() || role == "admin" {
		err := uco.userUseCase.DeleteById(idParam.String())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, helper.SuccesResponses("success delete user"))
	} else {
		// Jika bukan admin dan bukan pemilik data, kembalikan pesan "unauthorized".
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
	}
}

func (uco *UserController) GetAllUser(c echo.Context) error {
	// Mendapatkan role pengguna dari token
	role, err := middlewares.ExtractRole(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("error extracting role"))
	}

	if role != "admin" {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
	}

	// memanggil function dari usecase
	responseData, err := uco.userUseCase.FindAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error get data"))
	}

	userGetAllData := make([]UserResponse, 0)
	for _, value := range responseData {
		userResponse := MainResponse(value)
		userGetAllData = append(userGetAllData, userResponse)
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get all user data", userGetAllData))
}

func (uco *UserController) VerifyAccount(c echo.Context) error {
	token := c.QueryParam("token")

	user, err := uco.userUseCase.GetByVerificationToken(token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid or expired verification token"))
	}

	emailDone, _ := email.ParseTemplate("verification_active.html", nil)
	if user.IsVerified {
		return c.HTML(http.StatusOK, emailDone)
	}

	user.IsVerified = true

	err = uco.userUseCase.UpdateIsVerified(user.Id, true)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to activate the user"))
	}

	emailContent, _ := email.ParseTemplate("success_verification.html", nil)
    return c.HTML(http.StatusOK, emailContent)
}
