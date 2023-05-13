package users

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID        uint
	Name      string `validate:"required"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserService interface {
	Login(email string, password string) (Core, string, error)
	Create(input Core) error
}

type UserData interface {
	Login(email string) (Core, error)
	Insert(input Core) error
}
type UserDelivery interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
}
