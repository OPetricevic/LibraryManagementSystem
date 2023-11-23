package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"` // Use UUID as primary key
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Email    string `gorm:"unique"`
	Password string
	APIKey   string `gorm:"unique"` // Ensure APIKey is unique
}
