package data

import (
	"go-clean-aws/features/users"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

// Insert implements users.UserData
func (uq *userQuery) Insert(input users.Core) error {
	panic("unimplemented")
}

// Login implements users.UserData
func (uq *userQuery) Login(email string) (users.Core, error) {
	panic("unimplemented")
}

func New(db *gorm.DB) users.UserData {
	return &userQuery{
		db: db,
	}
}
