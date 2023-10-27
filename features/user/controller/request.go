package controller

import "recycle/features/user/entity"

type UserRequest struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Password    string  `json:"password"`
	Gender      string  `json:"gender"`
	Age         int     `json:"age"`
	Address     string  `json:"address"`
	SaldoPoints int `json:"saldo_points"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RequestMain(dataRequest UserRequest) entity.Main {
	return entity.Main{
		Id:          dataRequest.ID,
		Name:        dataRequest.Name,
		Email:       dataRequest.Email,
		Password:    dataRequest.Password,
		Gender:      dataRequest.Gender,
		Age:         dataRequest.Age,
		Address:     dataRequest.Address,
		SaldoPoints: dataRequest.SaldoPoints,
	}
}
