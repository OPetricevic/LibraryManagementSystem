package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title"`
	Author      string `json:"author"`
	ISBN        string `gorm:"unique;not null" json:"isbn"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Status      string `gorm:"not null;default:'available'" json:"status"`
	Quantity    int    `json:"quantity"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type BookCopy struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	BookID         uint       `json:"book_id"`                                    // Foreign key to the Book model
	Status         string     `gorm:"not null;default:'available'" json:"status"` // e.g., Available, Checked Out, Reserved
	Location       string     `json:"location"`                                   // Optional: location in the library
	Condition      string     `json:"condition"`                                  // e.g., New, Good, Worn
	CheckedOutDate *time.Time `json:"checked_out_date,omitempty"`
	DueDate        *time.Time `json:"due_date,omitempty"`
	ReturnedDate   *time.Time `json:"returned_date,omitempty"`
	gorm.Model
}
