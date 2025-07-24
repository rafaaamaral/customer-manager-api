package database

import (
	"customer-manager-api/src/config"
	"customer-manager-api/src/models"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {

	db, err := gorm.Open(sqlserver.Open(config.StringConnection), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migração automática
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Erro ao fazer a migração:", err)
	}

	return db, nil
}
