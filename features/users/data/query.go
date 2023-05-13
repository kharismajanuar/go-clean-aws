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
	dataModel := CoreToUserModel(input)
	txInsert := uq.db.Create(&dataModel)
	if txInsert.Error != nil {
		return txInsert.Error
	}
	if txInsert.RowsAffected == 0 {
		return txInsert.Error
	}
	return nil
}

// Login implements users.UserData
func (uq *userQuery) Login(email string) (users.Core, error) {
	dataModel := User{}
	txSelect := uq.db.Where("email = ?", email).First(&dataModel)
	if txSelect.Error != nil {
		return users.Core{}, txSelect.Error
	}
	return (UserModelToCore(dataModel)), nil
}

func New(db *gorm.DB) users.UserData {
	return &userQuery{
		db: db,
	}
}
