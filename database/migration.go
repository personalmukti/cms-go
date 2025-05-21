package database

import (
	"cms-go-2/config"
	"cms-go-2/models"
	"cms-go-2/seeders"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Get("DB_HOST"),
		config.Get("DB_PORT"),
		config.Get("DB_USER"),
		config.Get("DB_PASSWORD"),
		config.Get("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db

	// Migrasi tabel
	err = db.AutoMigrate(&models.Role{}, &models.User{})
	if err != nil {
		return err
	}

	// Panggil seeder terpisah
	seeders.SeedRoles(db)
	seeders.SeedUsers(db)

	return nil
}
