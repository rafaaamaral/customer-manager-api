package repositories

import (
	"customer-manager-api/src/models"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type users struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *users {
	return &users{db: db}
}

func (u users) CreateUser(user models.User) (uint, error) {
	user.UniqueCode = uuid.New() // Ensure unique code is set

	result := u.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	return user.ID, nil
}

func (u users) GetUsers(name string) ([]models.User, error) {
	var users []models.User
	result := u.db.Select("name, email, phone").Where("name like ?", fmt.Sprintf("%%%s%%", name)).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (u users) GetUserById(id uint) (models.User, error) {
	var user models.User
	result := u.db.Select("name, email, phone").First(&user, id)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func (u users) UpdateUser(user models.User) error {
	result := u.db.Model(&user).Updates(models.User{Name: user.Name, Email: user.Email, Phone: user.Phone})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u users) DeleteUser(id uint) error {
	result := u.db.Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("user with ID %d not found", id)
	}

	return nil
}
