package handlers

import (
	"errors"
	"unicode"
)

func validatePassword(password string) error {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasNumber  = false
		hasSpecial = false
	)

	if len(password) >= 8 {
		hasMinLen = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	if hasMinLen && hasUpper && hasNumber && hasSpecial {
		return nil
	}

	return errors.New("password must be at least 8 characters long, including an uppercase letter, a number, and a special character")
}
