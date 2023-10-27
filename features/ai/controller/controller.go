package controller

import (
	"net/http"
	"recycle/features/ai/entity"

	"github.com/labstack/echo/v4"
)

type RubbishController struct {
	rubbishUseCase entity.UseCaseInterface
}

func NewRubbishController(uc entity.UseCaseInterface) *RubbishController {
	return &RubbishController{
		rubbishUseCase: uc,
	}
}

func (uc *RubbishController) GetRecyclableRecommendation(c echo.Context) error {
	var requestBody RubbishRequest
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	if requestBody.Type == "" {
		return c.JSON(http.StatusBadRequest, "type field is required")
	}

	result, err := uc.rubbishUseCase.RecommendRecyclable(requestBody.Type)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to get data")
	}

	response := RubbishResponse{
		Status: "success",
		Result: result,
	}

	return c.JSON(http.StatusOK, response)
}
