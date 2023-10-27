package controller

import (
	"io/ioutil"
	"net/http"
	"recycle/app/middlewares"
	"recycle/email"
	"recycle/features/detail_pickup/entity"
	pick "recycle/features/pickup/entity"
	user "recycle/features/user/entity"
	"recycle/helper"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type DetailPickupController struct {
	detailPickupUseCase entity.UseCaseInterface
	pickupUseCase       pick.UseCaseInterface
	userUseCase         user.UseCaseInterface
}

func NewPickupControllers(uc entity.UseCaseInterface, pickupUseCase pick.UseCaseInterface, userUseCase user.UseCaseInterface) *DetailPickupController {
	return &DetailPickupController{
		detailPickupUseCase: uc,
		pickupUseCase:       pickupUseCase,
		userUseCase:         userUseCase,
	}
}

func (uco *DetailPickupController) CreateDetailPickup(c echo.Context) error {
	// Mendapatkan role pengguna dari token
	role, err := middlewares.ExtractRole(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("error extracting role"))
	}

	// Periksa apakah pengguna adalah admin.
	if role == "admin" {
		// Bind data
		dataInput := DetailPickupRequest{}
		errBind := c.Bind(&dataInput)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("error bind data"))
		}

		data := RequestDetailPickup(dataInput)

		totalUserPoints, errCreate := uco.detailPickupUseCase.Create(data)
		if errCreate != nil {
			if strings.Contains(errCreate.Error(), "validation") {
				return c.JSON(http.StatusBadRequest, helper.ErrorResponse(errCreate.Error()))
			} else {
				return c.JSON(http.StatusBadRequest, helper.ErrorResponse("failed to create data: "+errCreate.Error()))
			}
		}

		pickupData, err := uco.pickupUseCase.GetById(dataInput.PickupID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to get pickup data"))
		}

		emailUser, err := uco.userUseCase.GetEmailByID(pickupData.UserId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to get user data"))
		}

		userID := pickupData.UserId
		pickupDateStr := pickupData.PickupDate
		status := pickupData.Status
	
		pickupDate, err := time.Parse(time.RFC3339, pickupDateStr)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to parse pickup date"))
		}

		formattedDate := pickupDate.Format("2006-01-02")

		emailData := struct {
			UserID      string
			PickupDate  string
			PickupID    string
			Status      string
			TotalPoints int
		}{
			UserID:      userID,
			PickupDate:  formattedDate,
			PickupID:    dataInput.PickupID,
			Status:      status,
			TotalPoints: totalUserPoints,
		}

		emailTemplate, err := ioutil.ReadFile("email/templates/recycling_info.html")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to read email template"))
		}

		_, errEmail := email.SendEmailForPickup(emailUser, string(emailTemplate), emailData)
		if errEmail != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to send verification email"))
		}

		return c.JSON(http.StatusCreated, helper.SuccesResponses("success create data"))
	} else {
		// Jika pengguna bukan admin, kembalikan pesan "unauthorized".
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
	}
}

func (uco *DetailPickupController) GetAllDetailPickup(c echo.Context) error {
	role, err := middlewares.ExtractRole(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("error extracting role"))
	}

	if role == "admin" {
		// memanggil function dari usecase
		responseData, err := uco.detailPickupUseCase.FindAllDetailPickup()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error get data"))
		}

		DetailPickupData := MapModelsToController(responseData)

		return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get all detail pickup", DetailPickupData))
	} else {
		// Jika bukan admin, kembalikan pesan "unauthorized".
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
	}
}
