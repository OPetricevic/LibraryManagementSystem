package repository

import (
	"errors"

	models "github.com/OPetricevic/LibraryManagementSystem/Models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.Db.Create(user).Error
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := r.Db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// User not found
			return nil, errors.New("user not found")
		}
		// Other database error
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) GetAllUsers() ([]models.AllUserInformation, error) {
	var users []models.AllUserInformation
	result := r.Db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *UserRepository) UpdateUserByID(userID string, updateData models.UserUpdateRequest) error {

	updates := make(map[string]interface{})

	if updateData.FirstName != "" {
		updates["first_name"] = updateData.FirstName
	}
	if updateData.LastName != "" {
		updates["last_name"] = updateData.LastName
	}
	if updateData.Email != "" {
		updates["email"] = updateData.Email
	}
	if updateData.Role != "" {
		updates["role"] = updateData.Role
	}
	if updateData.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updateData.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		updates["password"] = string(hashedPassword)
	}

	return r.Db.Model(&models.User{}).Where("id = ?", userID).Updates(updates).Error
}
