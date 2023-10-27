package controller

import (
	"net/http"
	"recycle/app/middlewares"
	"recycle/features/rubbish/entity"
	"recycle/helper"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type RubbishController struct {
	rubbishUseCase entity.UseCaseInterface
}

func NewRubbishControllers(uc entity.UseCaseInterface) *RubbishController {
	return &RubbishController{
		rubbishUseCase: uc,
	}
}

func (uco *RubbishController) CreateRubbish(c echo.Context) error {
    // Mendapatkan role pengguna dari token
    role, err := middlewares.ExtractRole(c)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("error extracting role"))
    }

    // Periksa apakah pengguna adalah admin.
    if role == "admin" {
        // Bind data
        dataInput := RubbishRequest{}
        errBind := c.Bind(&dataInput)
        if errBind != nil {
            return c.JSON(http.StatusBadRequest, helper.ErrorResponse("error bind data"))
        }

        data := RequestMain(dataInput)

        image, err := c.FormFile("image_url")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("error uploading image"+err.Error()))
		}

        errCreate := uco.rubbishUseCase.Create(data, image)
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


func (uco *RubbishController) GetRubbish(c echo.Context) error {
	// Extra token dari id
	idToken := middlewares.ExtractToken(c)
	if idToken == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
	}

	idParamStr := c.Param("id")
	idParam, errId := uuid.Parse(idParamStr)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid UUID format"))
	}

	// Izinkan akses ke data sampah (rubbish) berdasarkan ID tanpa pemeriksaan tambahan.
	result, err := uco.rubbishUseCase.GetById(idParam.String())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error reading data"))
	}

	var rubbishResponse = MainResponse(result)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get rubbish", rubbishResponse))
}

func (uco *RubbishController) UpdateRubbish(c echo.Context) error {
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
            return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid UUID format"))
        }

        rubbishReq := RubbishRequest{}
        errBind := c.Bind(&rubbishReq)
        if errBind != nil {
            return c.JSON(http.StatusBadRequest, helper.ErrorResponse("error bind data"))
        }

        userMain := RequestMain(rubbishReq)

        image, err := c.FormFile("image_url")
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("error uploading image"+ err.Error()))
		}

        data, err := uco.rubbishUseCase.UpdateById(idParam.String(), userMain, image)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
        }

        return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success update data", MainResponse(data)))
    } else {
        // Jika bukan admin, kembalikan pesan "unauthorized".
        return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
    }
}

func (uco *RubbishController) DeleteRubbish(c echo.Context) error {
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
            return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid UUID format"))
        }

        err := uco.rubbishUseCase.DeleteById(idParam.String())
        if err != nil {
            return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
        }
        return c.JSON(http.StatusOK, helper.SuccesResponses("success delete rubbish"))
    } else {
        // Jika bukan admin, kembalikan pesan "unauthorized".
        return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
    }
}

func (uco *RubbishController) GetAllRubbish(c echo.Context) error {
    // Extra token dari id
	idToken := middlewares.ExtractToken(c)
	if idToken == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
	}

	// memanggil function dari usecase
	responseData, err := uco.rubbishUseCase.FindAllRubbish()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error get data"))
	}

	rubbishGetAllData := make([]RubbishRespon, 0)
	for _, value := range responseData {
		userResponse := MainResponse(value)
		rubbishGetAllData = append(rubbishGetAllData, userResponse)
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success get all user rubbish", rubbishGetAllData))
}

func (uco *RubbishController) GetRubbishByType(c echo.Context) error {
    // Extra token dari id
    idToken := middlewares.ExtractToken(c)
	if idToken == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("unauthorized"))
	}

    
    typeRubbish := c.QueryParam("type")
    result, err := uco.rubbishUseCase.GetByType(typeRubbish)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error reading location data"))
    }

    var rubbishResponse []RubbishRespon

    // Konversi setiap lokasi ke MainResponse
    for _, rubbish := range result {
        rubbishResponse = append(rubbishResponse, MainResponse(rubbish))
    }

    return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get rubbish", rubbishResponse))
}

