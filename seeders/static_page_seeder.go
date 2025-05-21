package seeders

import (
	"cms-go-2/models"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedStaticPages(db *gorm.DB) {
	var count int64
	db.Model(&models.StaticPage{}).Count(&count)

	if count == 0 {
		pages := []models.StaticPage{
			{
				ID:      uuid.New(),
				Title:   "Tentang Kami",
				Slug:    "about",
				Type:    "about",
				Content: "",
				Status:  "published",
			},
			{
				ID:      uuid.New(),
				Title:   "Kontak",
				Slug:    "contact",
				Type:    "contact",
				Content: "",
				Status:  "published",
			},
		}

		if err := db.Create(&pages).Error; err != nil {
			log.Println("Gagal seeding static pages:", err)
		} else {
			log.Println("Seeder static pages berhasil")
		}
	}
}
