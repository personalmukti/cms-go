package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(50);uniqueIndex;not null" json:"name"` // admin, editor, operator
	Label     string         `gorm:"type:varchar(100)" json:"label"`                    // Nama tampil, misal: "Admin"
	Desc      string         `gorm:"type:text" json:"desc"`                             // Deskripsi peran
	Users     []User         `gorm:"foreignKey:RoleID" json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (r *Role) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.New()
	return
}
