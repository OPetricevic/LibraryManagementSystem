package handlers

import (
	"encoding/json"
	"net/http"

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
	//Here I'm defining with govalidator if the user information that we are reading in the correct format that is defined in models.go
	if _, err := govalidator.ValidateStruct(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to create account", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	if err := uc.Repo.CreateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
