package models

import (
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string         `gorm:"type:char(36);primary_key;" json:"id" valid:"-"`
	CreatedAt time.Time      `json:"created_at" valid:"-"`
	UpdatedAt time.Time      `json:"updated_at" valid:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at" valid:"-"`

	FirstName string `gorm:"type:varchar(50);not null" json:"first_name" valid:"required,alpha"`
	LastName  string `gorm:"type:varchar(50);not null" json:"last_name" valid:"required,alpha"`
	Email     string `gorm:"type:varchar(255);unique;index;not null" json:"email" valid:"required,email"`
	Password  string `gorm:"not null" json:"password" valid:"required"`
	APIKey    string `gorm:"type:varchar(64);unique;index" json:"api_key" valid:"-"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// BeforeCreate pravi UUID
func (u *User) BeforeCreate(tx *gorm.DB) error {
	var err error
	u.ID, err = generateUUID()
	if err != nil {
		log.Printf("Error generating UUID: %v", err)
		return errors.New("failed to generate a unique identifier")
	}

	u.APIKey, err = generateAPIKey()
	if err != nil {
		log.Printf("Error generating API key: %v", err)
		return errors.New("failed to generate an API key")
	}
	return nil
}
