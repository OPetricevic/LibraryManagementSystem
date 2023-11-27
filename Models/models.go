package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"type:char(36);primary_key;"` // Changed from uuid.UUID to string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	FirstName string `gorm:"type:varchar(50)"`
	LastName  string `gorm:"type:varchar(50)"`
	Email     string `gorm:"type:varchar(255);unique;index"`
	Password  string
	APIKey    string `gorm:"type:varchar(64);unique;index"`
}

// BeforeCreate pravi UUID
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}
