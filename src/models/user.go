package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Birthdate time.Time `json:"birthdate,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	Password  string    `json:"password,omitempty"`
}
