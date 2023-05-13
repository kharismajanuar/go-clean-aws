package service

import (
	"errors"
	"go-clean-aws/features/users"
	"go-clean-aws/utils/helper"

	"github.com/go-playground/validator/v10"
)

type userUseCase struct {
	qry      users.UserData
	validate *validator.Validate
}

// Create implements users.UserService
func (uuc *userUseCase) Create(input users.Core) error {
	err := uuc.validate.Struct(input)
	if err != nil {
		return errors.New("validate: " + err.Error())
	}

	// hash password
	hashed, err := helper.HashPassword(input.Password)
	if err != nil {
		return err
	}
	input.Password = string(hashed)

	err = uuc.qry.Insert(input)
	if err != nil {
		return err
	}

	return nil
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
