package models

import "github.com/google/uuid"

type ArticleTag struct {
	ArticleID uuid.UUID `gorm:"type:uuid;primaryKey"`
	TagID     uuid.UUID `gorm:"type:uuid;primaryKey"`
}
