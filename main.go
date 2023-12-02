package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	handlers "github.com/OPetricevic/LibraryManagementSystem/Handlers"
	middleware "github.com/OPetricevic/LibraryManagementSystem/Middleware"
	models "github.com/OPetricevic/LibraryManagementSystem/Models"
	repository "github.com/OPetricevic/LibraryManagementSystem/Repository"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var jwtSecretKey []byte

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}
	jwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	if len(jwtSecretKey) == 0 {
		log.Fatal("JWT secret key must be set")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Category{}, &models.Book{}, &models.BookCopy{})
	if err != nil {
		log.Fatal("Failed to automigrate", err)
	}

	ensureAdminUserExists(db)

	userRepo := repository.NewUserRepository(db)
	bookRepo := repository.NewBookRepository(db)
	userController := handlers.NewUserController(userRepo, jwtSecretKey)
	adminController := handlers.NewAdminController(userRepo, bookRepo)

	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/register", userController.Register).Methods("POST")
	r.HandleFunc("/login", userController.Login).Methods("POST")

	// User routes
	userRoutes := r.PathPrefix("/").Subrouter()
	userRoutes.Use(middleware.JWTUserAuthMiddleware(jwtSecretKey))

	// Admin routes
	adminRoute := r.PathPrefix("/admin").Subrouter()
	adminRoute.Use(middleware.JWTAdminAuthMiddleware(jwtSecretKey))

	adminRoute.HandleFunc("/users", adminController.GetAllUsers).Methods("GET")            //admin/users
	adminRoute.HandleFunc("/users/{email}", adminController.GetUserByEmail).Methods("GET") // admin/users/dummyemail

	adminRoute.HandleFunc("/users/{id}", adminController.UpdateUser).Methods("PUT", "PATCH") //admin/users/id
	adminRoute.HandleFunc("/add_books", adminController.AddBook).Methods("POST")             //admin/addBooks
	adminRoute.HandleFunc("/books", adminController.ListBooks).Methods("GET")

	fmt.Printf("Server is running")
	err = http.ListenAndServe(":6666", r)
	if err != nil {
		log.Fatal("Failed to start server", err)
	}
}

func ensureAdminUserExists(db *gorm.DB) {
	defaultAdminEmail := os.Getenv("DEFAULT_ADMIN_EMAIL")
	defaultAdminPassword := os.Getenv("DEFAULT_ADMIN_PASSWORD")
	defaultAdminName := os.Getenv("DEFAULT_ADMIN_NAME")

	//Checks if an Admin acc exists
	var count int64
	db.Model(&models.User{}).Where("role = ?", "admin").Count(&count)
	if count == 0 {
		//Admin doesn't exist so we create one
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(defaultAdminPassword), bcrypt.DefaultCost)
		adminUser := models.User{
			Email:     defaultAdminEmail,
			Password:  string(hashedPassword),
			Role:      "admin",
			FirstName: defaultAdminName,
			LastName:  defaultAdminName,
		}
		db.Create(&adminUser)

	}
}
