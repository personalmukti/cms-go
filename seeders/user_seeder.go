package seeders

import (
	"cms-go-2/models"
	"cms-go-2/utils"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {
	var count int64
	db.Model(&models.User{}).Count(&count)

	if count > 0 {
		log.Println("User sudah tersedia, tidak perlu seeding ulang.")
		return
	}

	// Cari semua role yang diperlukan
	var adminRole, editorRole, operatorRole models.Role

	if err := db.Where("name = ?", "admin").First(&adminRole).Error; err != nil {
		log.Println("Gagal menemukan role admin:", err)
		return
	}
	if err := db.Where("name = ?", "editor").First(&editorRole).Error; err != nil {
		log.Println("Gagal menemukan role editor:", err)
		return
	}
	if err := db.Where("name = ?", "operator").First(&operatorRole).Error; err != nil {
		log.Println("Gagal menemukan role operator:", err)
		return
	}

	// Hash password sekali saja
	password := "password123"
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		log.Println("Gagal hash password:", err)
		return
	}

	users := []models.User{
		{
			ID:       uuid.New(),
			Name:     "Admin CMS",
			Email:    "admin@cms.go",
			Password: hashedPassword,
			RoleID:   adminRole.ID,
		},
		{
			ID:       uuid.New(),
			Name:     "Editor CMS",
			Email:    "editor@cms.go",
			Password: hashedPassword,
			RoleID:   editorRole.ID,
		},
		{
			ID:       uuid.New(),
			Name:     "Operator CMS",
			Email:    "operator@cms.go",
			Password: hashedPassword,
			RoleID:   operatorRole.ID,
		},
	}

	if err := db.Create(&users).Error; err != nil {
		log.Println("Gagal seeding users:", err)
	} else {
		log.Println("Seeder 3 user berhasil ditambahkan:")
		log.Println("📌 admin@cms.go / password123")
		log.Println("📌 editor@cms.go / password123")
		log.Println("📌 operator@cms.go / password123")
	}
}
