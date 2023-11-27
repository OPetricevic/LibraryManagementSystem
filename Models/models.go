package models

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
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
func (u *User) BeforeCreate(tx *gorm.DB) error {
	var err error
	u.ID, err = generateUUID()
	if err != nil {
		// Log the error and return a generic error message
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

func generateUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func generateAPIKey() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
