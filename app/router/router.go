package router

import (
	_userData "go-clean-aws/features/users/data"
	_userHandler "go-clean-aws/features/users/handler"
	_userService "go-clean-aws/features/users/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userData := _userData.New(db)
	userService := _userService.New(userData)
	userHandler := _userHandler.New(userService)

	e.POST("/register", userHandler.Register)
	e.POST("/login", userHandler.Login)
}
