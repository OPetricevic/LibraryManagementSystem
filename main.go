package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		log.Fatal("The DB_PASSWORD environment variable is not set.")
	}
	dsn := "root:" + dbPassword + "@tcp(localhost:3306)/Library_db"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Make sure the connection is available.
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database is reachable!")
}
