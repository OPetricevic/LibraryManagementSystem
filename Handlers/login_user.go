package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	models "github.com/OPetricevic/LibraryManagementSystem/Models"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type CustomClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := uc.Repo.GetUserByEmail(loginRequest.Email)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	tokenString, err := uc.generateJWTToken(user.ID)
	if err != nil {
		http.Error(w, "Failed to generate JWT token", http.StatusInternalServerError)
		return
	}

	response := struct {
		Token string `json:"token"`
	}{
		Token: tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (uc *UserController) generateJWTToken(userID string) (string, error) {
	// Set up claims
	claims := CustomClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
			Issuer:    "LibraryManagementSystem",
			IssuedAt:  time.Now().Unix(),
		},
	}

	// Create a new token object with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with your secret key
	tokenString, err := token.SignedString(uc.jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
