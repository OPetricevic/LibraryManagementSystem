package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	models "github.com/OPetricevic/LibraryManagementSystem/Models"
	"gorm.io/gorm"
)

func (ac *adminController) AddBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	if err := validateBookData(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if book.Category.Name != "" {
		var category models.Category
		result := ac.BookRepo.Db.First(&category, "name = ?", book.Category.Name)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			category = models.Category{Name: book.Category.Name}
			ac.BookRepo.Db.Create(&category)
		}
		book.CategoryID = category.ID
	} else {
		book.CategoryID = ac.BookRepo.GetOrCreateUncategorizedCategory().ID
	}

	if err := ac.BookRepo.AddBook(&book); err != nil {
		http.Error(w, "Failed to add book: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var addCopyErrors []string
	for i := 0; i < book.Quantity; i++ {
		bookCopy := models.BookCopy{BookID: book.ID, Status: "available"}
		if err := ac.BookRepo.AddBookCopy(&bookCopy); err != nil {
			addCopyErrors = append(addCopyErrors, err.Error())
		}
	}

	if len(addCopyErrors) > 0 {
		errorMsg := "Invalid book data:\n" + strings.Join(addCopyErrors, ";\n") + ";\n"
		http.Error(w, errorMsg, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func (ac *adminController) ListBooks(w http.ResponseWriter, r *http.Request) {

}
