package repositories

import (
	"customer-manager-api/src/models"

	"gorm.io/gorm"
)

type users struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *users {
	return &users{db: db}
}

func (u users) CreateUser(user models.User) (uint64, error) {
	result := u.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	return uint64(user.ID), nil
}
