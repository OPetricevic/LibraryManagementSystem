package handlers

import repository "github.com/OPetricevic/LibraryManagementSystem/Repository"

type UserController struct {
	Repo *repository.UserRepository
}

func NewUserController(repo *repository.UserRepository) *UserController {
	return &UserController{Repo: repo}
}
