package database

import (
	"customer-manager-api/src/config"
	"customer-manager-api/src/models"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	DB, err := gorm.Open(sqlserver.Open(config.StringConnection), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Migração automática
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Erro ao fazer a migração:", err)
	}
}
