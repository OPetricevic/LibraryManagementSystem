package repository

import (
	models "github.com/OPetricevic/LibraryManagementSystem/Models"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func newUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.Db.Create(user).Error
}
