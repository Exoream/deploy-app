package controller

import (
	"net/http"
	"recycle/app/middlewares"
	"recycle/features/location/entity"
	"recycle/helper"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type LocationController struct {
	locationUseCase entity.UseCaseInterface
}

func NewLocationControllers(uc entity.UseCaseInterface) *LocationController {
	return &LocationController{
		locationUseCase: uc,
	}
}

func (uco *LocationController) CreateLocation(c echo.Context) error {
	// Mendapatkan role pengguna dari token
	role, err := middlewares.ExtractRole(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("error extracting role"))
	}

	// Periksa apakah pengguna adalah admin.
	if role == "admin" {
		// Bind data
		dataInput := LocationRequest{}
		errBind := c.Bind(&dataInput)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("error bind data"))
		}

		data := RequestMain(dataInput)

		errCreate := uco.locationUseCase.Create(data)
		if errCreate != nil {
			if strings.Contains(errCreate.Error(), "validation") {
				return c.JSON(http.StatusBadRequest, helper.ErrorResponse(errCreate.Error()))
			} else {
				return c.JSON(http.StatusBadRequest, helper.ErrorResponse("failed to create data"))
			}
		}

		return c.JSON(http.StatusCreated, helper.SuccesResponses("success create data"))
	} else {
		// Jika pengguna bukan admin, kembalikan pesan "unauthorized".
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
	}
}

func (uco *LocationController) GetLocation(c echo.Context) error {
	// Extra token dari id
	idToken := middlewares.ExtractToken(c)
	if idToken == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
	}

	idParamStr := c.Param("id")
	idParam, errId := uuid.Parse(idParamStr)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid UUID format"))
	}

	result, err := uco.locationUseCase.GetById(idParam.String())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error reading location data"))
	}

	var locationResponse = MainResponse(result)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get location", locationResponse))
}

func (uco *LocationController) GetLocationByCity(c echo.Context) error {
    // Extra token dari id
    idToken := middlewares.ExtractToken(c)
	if idToken == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
	}

    
    city := c.QueryParam("city")
    result, err := uco.locationUseCase.GetByCity(city)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error reading location data"))
    }

    var locationResponse []LocationRespon

    // Konversi setiap lokasi ke MainResponse
    for _, location := range result {
        locationResponse = append(locationResponse, MainResponse(location))
    }

    return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get location", locationResponse))
}


func (uco *LocationController) UpdateLocation(c echo.Context) error {
    // Mendapatkan role pengguna dari token
    role, err := middlewares.ExtractRole(c)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("error extracting role"))
    }

    // Periksa apakah pengguna adalah admin.
    if role == "admin" {
        idParamStr := c.Param("id")
        idParam, errId := uuid.Parse(idParamStr)
        if errId != nil {
            return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid UUID format"))
        }

        locationReq := LocationRequest{}
        errBind := c.Bind(&locationReq)
        if errBind != nil {
            return c.JSON(http.StatusBadRequest, helper.ErrorResponse("error bind data"))
        }

        userMain := RequestMain(locationReq)
        data, err := uco.locationUseCase.UpdateById(idParam.String(), userMain)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
        }

        return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success update data", MainResponse(data)))
    } else {
        // Jika bukan admin, kembalikan pesan "unauthorized".
        return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
    }
}

func (uco *LocationController) DeleteLocation(c echo.Context) error {
    // Mendapatkan role pengguna dari token
    role, err := middlewares.ExtractRole(c)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("error extracting role"))
    }

    // Periksa apakah pengguna adalah admin.
    if role == "admin" {
        idParamStr := c.Param("id")
        idParam, errId := uuid.Parse(idParamStr)
        if errId != nil {
            return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid UUID format"))
        }

        err := uco.locationUseCase.DeleteById(idParam.String())
        if err != nil {
            return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
        }
        return c.JSON(http.StatusOK, helper.SuccesResponses("success delete location"))
    } else {
        // Jika bukan admin, kembalikan pesan "unauthorized".
        return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
    }
}

func (uco *LocationController) GetAllLocation(c echo.Context) error {
    // Extra token dari id
	idToken := middlewares.ExtractToken(c)
	if idToken == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
	}

	// memanggil function dari usecase
	responseData, err := uco.locationUseCase.FindAllLocation()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error get data"))
	}

	locationGetAllData := make([]LocationRespon, 0)
	for _, value := range responseData {
		userResponse := MainResponse(value)
		locationGetAllData = append(locationGetAllData, userResponse)
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get all location", locationGetAllData))
}
