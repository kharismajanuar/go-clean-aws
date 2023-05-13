package handler

import (
	"go-clean-aws/features/users"
	"net/http"

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
	input := RegisterRequest{}
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error bind data",
		})
	}
	err := uh.srv.Create(registerToCore(input))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"status":  "success",
		"message": "success register account",
	})
}

func New(srv users.UserService) users.UserDelivery {
	return &userHandler{
		srv: srv,
	}
}
