package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	handlers "github.com/OPetricevic/LibraryManagementSystem/Handlers"
	models "github.com/OPetricevic/LibraryManagementSystem/Models"
	repository "github.com/OPetricevic/LibraryManagementSystem/Repository"
	routes "github.com/OPetricevic/LibraryManagementSystem/Routes"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}
	fmt.Println("DB_USER:", os.Getenv("DB_USER"))
	fmt.Println("DB_PASSWORD:", os.Getenv("DB_PASSWORD"))
	fmt.Println("DB_NAME:", os.Getenv("DB_NAME"))
	fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))
	fmt.Println("DB_PORT:", os.Getenv("DB_PORT"))

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

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to automigrate", err)
	}

	userRepo := repository.NewUserRepository(db)
	userController := handlers.NewUserController(userRepo)

	mux := http.NewServeMux()
	routes.RegisterRoutes(mux, userController)

	fmt.Printf("Server is running")
	err = http.ListenAndServe(":6666", mux)
	if err != nil {
		log.Fatal("Failed to start server", err)
	}
}
