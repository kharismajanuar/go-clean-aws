package handler

import (
	"go-clean-aws/features/users"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	srv users.UserService
}

// Login implements users.UserDelivery
func (uh *userHandler) Login(c echo.Context) error {
	panic("unimplemented")
}

// Register implements users.UserDelivery
func (uh *userHandler) Register(c echo.Context) error {
	panic("unimplemented")
}

func New(srv users.UserService) users.UserDelivery {
	return &userHandler{
		srv: srv,
	}
}
