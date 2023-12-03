package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	models "github.com/OPetricevic/LibraryManagementSystem/Models"
)

func (ac *adminController) AddBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Find or create the category
	category, err := ac.BookRepo.FindOrCreateCategory(book.TempCategoryName)
	if err != nil {
		http.Error(w, "Failed to find or create category: "+err.Error(), http.StatusInternalServerError)
		return
	}
	book.CategoryID = category.ID

	// Add book to the database
	if err := ac.BookRepo.AddBook(&book); err != nil {
		http.Error(w, "Failed to add book: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Add book copies logic
	var addCopyErrors []string
	for i := 0; i < book.Quantity; i++ {
		bookCopy := models.BookCopy{BookID: book.ID, Status: "available"}
		if err := ac.BookRepo.AddBookCopy(&bookCopy); err != nil {
			addCopyErrors = append(addCopyErrors, err.Error())
		}
	}

	// Handle errors in adding book copies
	if len(addCopyErrors) > 0 {
		errorMsg := "Failed to add some book copies: " + strings.Join(addCopyErrors, "; ")
		http.Error(w, errorMsg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func (ac *adminController) ListBooks(w http.ResponseWriter, r *http.Request) {

}

func (ac *adminController) GetBookByCategory(w http.ResponseWriter, r *http.Request) {

}
