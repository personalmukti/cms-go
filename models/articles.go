package models

import (
	"cms-go-2/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Article struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Title     string         `gorm:"type:varchar(200);not null" json:"title"`
	Slug      string         `gorm:"type:varchar(255);uniqueIndex;not null" json:"slug"`
	Content   string         `gorm:"type:text;not null" json:"content"`
	Status    string         `gorm:"type:varchar(20);not null" json:"status"` // draft/published
	ImageURL  string         `gorm:"type:varchar(255)" json:"image_url"`
	AuthorID  uuid.UUID      `gorm:"type:uuid" json:"author_id"`
	Author    User           `gorm:"foreignKey:AuthorID" json:"author"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (a *Article) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()

	if a.Slug == "" {
		a.Slug = utils.GenerateUniqueSlug(tx, a.Title)
	}

	return
}
