package handlers

import (
	"encoding/json"
	"net/http"

	models "github.com/OPetricevic/LibraryManagementSystem/Models"
	repository "github.com/OPetricevic/LibraryManagementSystem/Repository"
	"github.com/gorilla/mux"
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
	// Validate the password if it's provided
if updateReq.Password != "" {
	if err := validatePassword(updateReq.Password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	

	if err := ac.Repo.UpdateUserByID(userID, updateData); err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully"})
}
