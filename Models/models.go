package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"` 
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	FirstName string `gorm:"type:varchar(50)"`
    LastName  string `gorm:"type:varchar(50)"`
    Email     string `gorm:"type:varchar(255);unique;index"`
    Password  string 
    APIKey    string `gorm:"type:varchar(64);unique;index"`
}
