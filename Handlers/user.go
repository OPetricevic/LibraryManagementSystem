package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"

	models "github.com/OPetricevic/LibraryManagementSystem/Models"
	repository "github.com/OPetricevic/LibraryManagementSystem/Repository"
)

type UserController struct {
	Repo *repository.UserRepository
}

func NewUserController(repo *repository.UserRepository) *UserController {
	return &UserController{Repo: repo}
}

func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	var errorMsgs []string // Declare the errorMsgs variable

	// Check if the user information is in the correct format defined in models.go
	if _, err := govalidator.ValidateStruct(user); err != nil {
		errorMsgs = append(errorMsgs, strings.Split(err.Error(), ";")...)
	}

	// Check if the password meets complexity requirements
	if err := validatePassword(user.Password); err != nil {
		errorMsgs = append(errorMsgs, "Invalid password")
	}

	if len(errorMsgs) > 0 {
		formattedMsg := strings.Join(errorMsgs, "\n")
		http.Error(w, formattedMsg, http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to create account", http.StatusInternalServerError)
		return
	}
	//Converting the password into a hashable string.
	user.Password = string(hashedPassword)

	if err := uc.Repo.CreateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response := map[string]string{"message": "Registered successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
