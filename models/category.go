package models

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"unique;not null" json:"name"`
	Slug      string    `gorm:"unique;not null" json:"slug"`
	CreatedAt time.Time `json:"created_at"`
}
