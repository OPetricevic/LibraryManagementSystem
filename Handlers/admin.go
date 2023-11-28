package handlers

import (
	"encoding/json"
	"net/http"

	repository "github.com/OPetricevic/LibraryManagementSystem/Repository"
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
