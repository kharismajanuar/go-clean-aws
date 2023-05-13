package handler

import "go-clean-aws/features/users"

type RegisterRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func registerToCore(data RegisterRequest) users.Core {
	return users.Core{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
}

func loginToCore(data LoginRequest) users.Core {
	return users.Core{
		Email:    data.Email,
		Password: data.Password,
	}
}
