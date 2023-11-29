package handlers

import (
	"errors"
	"strings"

	models "github.com/OPetricevic/LibraryManagementSystem/Models"
)

func validateBookData(book *models.Book) error {
	var validationErrors []string
	if book.Title == "" {
		validationErrors = append(validationErrors, "Title is required")
	}

	if book.ISBN == "" {
		validationErrors = append(validationErrors, "ISBN is required")
	}

	if book.Quantity <= 0 {
		validationErrors = append(validationErrors, "Quantity is required")
	}

	if len(validationErrors) > 0 {
		return errors.New(strings.Join(validationErrors, "; "))
	}

	return nil
}
