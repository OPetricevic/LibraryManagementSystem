package handlers

import (
	"encoding/json"
	"net/http"

	models "github.com/OPetricevic/LibraryManagementSystem/Models"
	repository "github.com/OPetricevic/LibraryManagementSystem/Repository"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type adminController struct {
	Repo *repository.UserRepository
}

func NewAdminController(repo *repository.UserRepository) *adminController {
	return &adminController{Repo: repo}
}

func (ac *adminController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := ac.Repo.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to get users", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (ac *adminController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	var updateData models.UserUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if updateData.Password != "" {
		if err := validatePassword(updateData.Password); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	//I has the password, if it was updated.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updateData.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	updateData.Password = string(hashedPassword)

	if err := ac.Repo.UpdateUserByID(userID, updateData); err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully"})
}
