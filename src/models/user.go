package models

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UniqueCode uuid.UUID `json:"unique_code,omitempty" gorm:"type:uniqueidentifier;uniqueIndex;not null"`
	Name       string    `json:"name,omitempty" gorm:"type:varchar(100);not null"`
	Email      string    `json:"email,omitempty" gorm:"type:varchar(100);uniqueIndex;not null"`
	Birthdate  time.Time `json:"birthdate,omitempty" gorm:"type:date"`
	Phone      string    `json:"phone,omitempty" gorm:"type:varchar(15);"`
	Password   string    `json:"password,omitempty" gorm:"type:varchar(255);not null"`
}

func (user *User) PrepareToSave(isUpdate bool) error {
	if err := user.validate(isUpdate); err != nil {
		return err
	}

	user.formatData()

	return nil
}

func (user *User) validate(isUpdate bool) error {
	if user.Name == "" {
		return errors.New("name is required")
	}

	if user.Email == "" {
		return errors.New("email is required")
	}

	if !isUpdate && user.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

func (user *User) formatData() {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Phone = strings.TrimSpace(user.Phone)
}
