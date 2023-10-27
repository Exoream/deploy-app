package controller

import (
	// "recycle/features/pickup/controller"
	"recycle/features/pickup/controller"
	user "recycle/features/user/entity"

	// pickup "recycle/features/pickup/entity"
	"time"
)

type UserLoginResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type UserResponse struct {
	Id          string                    `json:"id"`
	Name        string                    `json:"name"`
	Email       string                    `json:"email"`
	Gender      string                    `json:"gender"`
	Age         int                       `json:"age"`
	Address     string                    `json:"address"`
	SaldoPoints int                   `json:"saldo_points"`
	CreatedAt   time.Time                 `json:"created_at"`
	Pickup      []controller.PickupRespon `json:"pickup"`
}

func MainResponse(dataMain user.Main) UserResponse {
	// Implementasi MainResponse sesuai dengan UserResponse
	userResponse := UserResponse{
		Id:          dataMain.Id,
		Name:        dataMain.Name,
		Email:       dataMain.Email,
		Gender:      dataMain.Gender,
		Age:         dataMain.Age,
		Address:     dataMain.Address,
		SaldoPoints: dataMain.SaldoPoints,
		CreatedAt:   dataMain.CreatedAt,
		Pickup:      []controller.PickupRespon{},
	}

	userResponse.Pickup = controller.MapModelToController(dataMain.Pickups)

	return userResponse
}

func LoginResponse(id, email, token string) UserLoginResponse {
	return UserLoginResponse{
		Id:    id,
		Email: email,
		Token: token,
	}
}
