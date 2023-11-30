package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title            string `gorm:"not null" json:"title"`
	Author           string `json:"author"`
	ISBN             string `gorm:"unique;not null" json:"isbn"`
	Description      string `json:"description"`
	CategoryID       uint   `gorm:"not null" json:"category_id"`
	TempCategoryName string `json:"category" gorm:"-"`
	Status           string `gorm:"not null;default:'available'" json:"status"`
	Quantity         int    `gorm:"not null" json:"quantity"`
}

type Category struct {
	gorm.Model
	Name        string `gorm:"unique;not null" json:"name"`
	Description string `json:"description"`
	Books       []Book `gorm:"foreignKey:CategoryID" json:"books"`
}

type BookCopy struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	BookID         uint       `json:"book_id"`
	Book           Book       `json:"book"`
	Status         string     `gorm:"not null;default:'available'" json:"status"`
	CheckedOutDate *time.Time `json:"checked_out_date,omitempty"`
	DueDate        *time.Time `json:"due_date,omitempty"`
	ReturnedDate   *time.Time `json:"returned_date,omitempty"`
	gorm.Model
}
