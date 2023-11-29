package repository

import (
	models "github.com/OPetricevic/LibraryManagementSystem/Models"
	"gorm.io/gorm"
)

type BookRepository struct {
	Db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{Db: db}
}

func (r *BookRepository) AddBook(book *models.Book) error {
	return r.Db.Create(book).Error
}

func (r *BookRepository) AddBookCopy(bookCopy *models.BookCopy) error {
	return r.Db.Create(bookCopy).Error
}

func (r *BookRepository) ListBooks() ([]models.Book, error) {
	var books []models.Book
	result := r.Db.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}
