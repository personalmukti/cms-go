package seeders

import (
	"cms-go-2/models"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedArticles(db *gorm.DB) {
	var count int64
	db.Model(&models.Article{}).Count(&count)

	if count > 0 {
		log.Println("Artikel sudah tersedia, tidak perlu seeding ulang.")
		return
	}

	// Coba cari penulis: editor@cms.go atau admin@cms.go
	var author models.User
	if err := db.Where("email = ?", "editor@cms.go").First(&author).Error; err != nil {
		if err := db.Where("email = ?", "admin@cms.go").First(&author).Error; err != nil {
			log.Println("Gagal menemukan user editor/admin untuk author artikel.")
			return
		}
	}

	articles := []models.Article{
		{
			ID:       uuid.New(),
			Title:    "Artikel Dummy 1",
			Content:  "Ini adalah isi artikel dummy pertama.",
			Status:   "published",
			ImageURL: "https://placehold.co/600x400",
			AuthorID: author.ID,
		},
		{
			ID:       uuid.New(),
			Title:    "Artikel Dummy 2",
			Content:  "Ini adalah isi artikel dummy kedua.",
			Status:   "published",
			ImageURL: "https://placehold.co/600x400",
			AuthorID: author.ID,
		},
	}

	if err := db.Create(&articles).Error; err != nil {
		log.Println("Gagal seeding artikel:", err)
	} else {
		log.Println("Seeder artikel berhasil ditambahkan.")
	}
}
