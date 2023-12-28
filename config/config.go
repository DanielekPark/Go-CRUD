package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var db *sql.DB

func Config() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
	}

	// Database connection
	db, err = sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatal("failed to open db connection", err)
	}
}
