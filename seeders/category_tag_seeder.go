package seeders

import (
	"cms-go-2/models"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedCategoryAndTag(db *gorm.DB) {
	// Cek kategori
	var count int64
	db.Model(&models.Category{}).Count(&count)
	if count == 0 {
		categories := []models.Category{
			{ID: uuid.New(), Name: "Teknologi", Slug: "teknologi"},
			{ID: uuid.New(), Name: "Gaya Hidup", Slug: "gaya-hidup"},
			{ID: uuid.New(), Name: "Ekonomi", Slug: "ekonomi"},
		}
		db.Create(&categories)
		log.Println("Seeder kategori berhasil")
	}

	// Cek tag
	db.Model(&models.Tag{}).Count(&count)
	if count == 0 {
		tags := []models.Tag{
			{ID: uuid.New(), Name: "Tips", Slug: "tips"},
			{ID: uuid.New(), Name: "AI", Slug: "ai"},
			{ID: uuid.New(), Name: "Programming", Slug: "programming"},
		}
		db.Create(&tags)
		log.Println("Seeder tag berhasil")
	}
}
