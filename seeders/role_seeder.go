package seeders

import (
	"cms-go-2/models"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	var count int64
	db.Model(&models.Role{}).Count(&count)

	if count > 0 {
		log.Println("Role sudah tersedia, tidak perlu seeding ulang.")
		return
	}

	roles := []models.Role{
		{ID: uuid.New(), Name: "admin", Label: "Administrator", Desc: "Dapat mengelola seluruh sistem"},
		{ID: uuid.New(), Name: "editor", Label: "Editor", Desc: "Dapat mengelola artikel"},
		{ID: uuid.New(), Name: "operator", Label: "Operator", Desc: "Dapat mengelola halaman statis"},
	}

	if err := db.Create(&roles).Error; err != nil {
		log.Println("Gagal seeding roles:", err)
	} else {
		log.Println("Seeder role berhasil ditambahkan.")
	}
}
