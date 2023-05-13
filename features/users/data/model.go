package data

import (
	"go-clean-aws/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
}

func UserModelToCore(data User) users.Core {
	return users.Core{
		ID:       data.ID,
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
}

func CoreToUserModel(data users.Core) User {
	return User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
}
