package controller

import (
	"net/http"
	"recycle/app/middlewares"
	"recycle/features/pickup/entity"
	"recycle/helper"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PickupController struct {
	pickupUseCase entity.UseCaseInterface
}

func NewPickupControllers(uc entity.UseCaseInterface) *PickupController {
	return &PickupController{
		pickupUseCase: uc,
	}
}

func (uco *PickupController) CreatePickup(c echo.Context) error {
	// Mendapatkan role pengguna dari token
	role, err := middlewares.ExtractRole(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("error extracting role"))
	}

	// Mendapatkan user_id dari token
	idToken := middlewares.ExtractToken(c)
	if idToken == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
	}

	// Bind data
	dataInput := PickupRequest{}
	errBind := c.Bind(&dataInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("error bind data"))
	}

	// Konversi dataInput menjadi data Main
	data := RequestMain(dataInput)
	data.UserId = idToken.String()

	// Periksa apakah pengguna adalah "user" dan ID yang login sesuai dengan ID yang dikirim dalam permintaan
	if role == "user" && data.UserId == idToken.String() {
		image, err := c.FormFile("image_url")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("error uploading image"+err.Error()))
		}

		errCreate := uco.pickupUseCase.Create(data, image)
		if errCreate != nil {
			if strings.Contains(errCreate.Error(), "validation") {
				return c.JSON(http.StatusBadRequest, helper.ErrorResponse(errCreate.Error()))
			} else {
				return c.JSON(http.StatusBadRequest, helper.ErrorResponse("failed to create data:"+errCreate.Error()))
			}
		}

		return c.JSON(http.StatusCreated, helper.SuccesResponses("success create data"))
	} else {
		// Jika pengguna bukan user atau ID tidak sesuai, kembalikan pesan "unauthorized".
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
	}
}

func (uco *PickupController) UpdatePickup(c echo.Context) error {
	// Mendapatkan role pengguna dari token
	role, err := middlewares.ExtractRole(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("error extracting role"))
	}

	// Periksa apakah pengguna adalah user.
	if role == "user" {
		// Mendapatkan user_id dari token
		userID := middlewares.ExtractToken(c)
		if userID == uuid.Nil {
			return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
		}

		idParamStr := c.Param("id")
		idParam, errID := uuid.Parse(idParamStr)
		if errID != nil {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid UUID format"))
		}

		// Mengambil data pickup dari database berdasarkan ID pickup yang diberikan
		pickup, err := uco.pickupUseCase.GetById(idParam.String())
		if err != nil {
			return c.JSON(http.StatusNotFound, helper.ErrorResponse("pickup not found"))
		}

		// Periksa apakah pickup sesuai dengan pengguna yang sedang masuk
		if pickup.UserId != userID.String() {
			return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
		}

		pickupReq := PickupRequest{}
		errBind := c.Bind(&pickupReq)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("error bind data"))
		}
		pickupMain := RequestMain(pickupReq)

		image, err := c.FormFile("image_url")
		if err != nil {
			if err != http.ErrMissingFile {
				return c.JSON(http.StatusBadRequest, helper.ErrorResponse("error uploading image: " + err.Error()))
			}
		}

		data, err := uco.pickupUseCase.UpdateById(idParam.String(), pickupMain, image)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success update data", MainResponse(data)))
	} else {
		// Jika bukan user, kembalikan pesan "unauthorized".
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
	}
}

func (uco *PickupController) DeletePickup(c echo.Context) error {
	// Mendapatkan role pengguna dari token
	role, err := middlewares.ExtractRole(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("error extracting role"))
	}

	// Periksa apakah pengguna adalah user.
	if role == "user" {
		userID := middlewares.ExtractToken(c)
		if userID == uuid.Nil {
			return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
		}

		idParamStr := c.Param("id")
		idParam, errId := uuid.Parse(idParamStr)
		if errId != nil {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid UUID format"))
		}

		pickup, errData := uco.pickupUseCase.GetById(idParam.String())
		if errData != nil {
			return c.JSON(http.StatusNotFound, helper.ErrorResponse("pickup not found"))
		}

		// Periksa apakah pickup sesuai dengan pengguna yang sedang masuk
		if pickup.UserId != userID.String() {
			return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
		}

		err := uco.pickupUseCase.DeleteById(idParam.String())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.SuccesResponses("success delete pickup"))
	} else {
		// Jika bukan user, kembalikan pesan "unauthorized".
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
	}
}

func (uco *PickupController) GetAllPickup(c echo.Context) error {
	role, err := middlewares.ExtractRole(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("error extracting role"))
	}

	if role == "admin" {
		// memanggil function dari usecase
		responseData, err := uco.pickupUseCase.FindAllPickup()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error get data"))
		}

		PickupGetAllData := MapModelsToController(responseData)

		return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get all pickup", PickupGetAllData))
	} else {
		// Jika bukan admin, kembalikan pesan "unauthorized".
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
	}
}

func (uco *PickupController) GetDataByStatus(c echo.Context) error {
	role, err := middlewares.ExtractRole(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("error extracting role"))
	}

	if role == "admin" {
		status := c.QueryParam("status")
		result, err := uco.pickupUseCase.GetByStatus(status)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error reading data"))
		}

		var statusResponse []PickupResponForGetAll

		// Konversi setiap lokasi ke MainResponse
		for _, value := range result {
			statusResponse = append(statusResponse, MainResponses(value))
		}

		return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get data", statusResponse))
	} else {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
	}
}
