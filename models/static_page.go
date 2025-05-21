package models

import (
	"time"

	"github.com/google/uuid"
)

type StaticPage struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	Slug      string    `gorm:"uniqueIndex;not null" json:"slug"`
	Type      string    `gorm:"not null;uniqueIndex" json:"type"`
	Content   string    `gorm:"type:text" json:"content"`
	Status    string    `gorm:"type:varchar(20);default:'draft'" json:"status"` // draft / published
	UpdatedBy uuid.UUID `gorm:"type:uuid" json:"updated_by"`

	UpdatedAt time.Time `json:"updated_at"`
}
