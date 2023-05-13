package service

import (
	"errors"
	"go-clean-aws/features/users"
	"go-clean-aws/middlewares"
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
	hashed, err := helper.PassBcrypt(input.Password)
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
func (uuc *userUseCase) Login(input users.Core) (users.Core, string, error) {
	err := uuc.validate.StructExcept(input, "Name")
	if err != nil {
		return users.Core{}, "", errors.New("validate: " + err.Error())
	}

	// find account
	core, err := uuc.qry.Login(input.Email)
	if err != nil {
		return users.Core{}, "", err
	}

	err = helper.PassCompare(core.Password, input.Password)
	if err != nil {
		return users.Core{}, "", err
	}

	token, err := middlewares.CreateToken(core.ID)
	if err != nil {
		return users.Core{}, "", err
	}

	return core, token, nil
}

func New(ud users.UserData) users.UserService {
	return &userUseCase{
		qry:      ud,
		validate: validator.New(),
	}
}
