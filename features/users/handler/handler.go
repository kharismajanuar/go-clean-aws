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
	input := LoginRequest{}
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error bind data",
		})
	}

	core, token, err := uh.srv.Login(loginToCore(input))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success login",
		"id":      core.ID,
		"name":    core.Name,
		"token":   token,
	})
}

// Register implements users.UserDelivery
func (uh *userHandler) Register(c echo.Context) error {
	input := RegisterRequest{}
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error bind data",
		})
	}
	err = uh.srv.Create(registerToCore(input))
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
