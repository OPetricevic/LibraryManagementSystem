package repository

import (
	"errors"

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
func (r *BookRepository) GetOrCreateUncategorizedCategory() *models.Category {
	var category models.Category

	result := r.Db.First(&category, "name = ?", "uncategorized")
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		category = models.Category{Name: "uncategorized"}
		r.Db.Create(&category)
	}
	return &category
}

func (r *BookRepository) FindOrCreateCategory(categoryName string) (models.Category, error) {
	var category models.Category
	if result := r.Db.Where("name = ?", categoryName).FirstOrCreate(&category, models.Category{Name: categoryName}); result.Error != nil {
		return models.Category{}, result.Error
	}
	return category, nil
}

func (r *BookRepository) GetBooks(page int, pageSize int) ([]models.Book, error) {
	var books []models.Book
	offset := (page - 1) * pageSize
	result := r.Db.Preload("Category").Offset(offset).Limit(pageSize).Find(&books)
	//Adding Category to books, so I can list the categories with them.
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

// Gets All Books through Category
func (r *BookRepository) GetBookByCategory(categoryName string) ([]models.Book, error) {
	var books []models.Book
	result := r.Db.Preload("Category").Where("categories.name = ?", categoryName).Joins("JOIN categories ON categories.id = books.category_id").Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}
