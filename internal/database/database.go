package database

import (
	"fmt"
	"go-auth-rbac/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(dsn string) error {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Auto migrate models
	err = DB.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Permission{},
	)
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("Database connected successfully")
	return nil
}
