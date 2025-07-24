package models

import (
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
