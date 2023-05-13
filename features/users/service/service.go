package service

import (
	"go-clean-aws/features/users"

	"github.com/go-playground/validator/v10"
)

type userUseCase struct {
	qry      users.UserData
	validate *validator.Validate
}

// Create implements users.UserService
func (uuc *userUseCase) Create(input users.Core) error {
	panic("unimplemented")
}

// Login implements users.UserService
func (uuc *userUseCase) Login(email string, password string) (users.Core, string, error) {
	panic("unimplemented")
}

func New(ud users.UserData) users.UserService {
	return &userUseCase{
		qry:      ud,
		validate: validator.New(),
	}
}
