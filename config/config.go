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

	// Open a connection to the database
	db, err = sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatal("failed to open db connection", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping: %v line 19", err)
	}
	log.Println("Successfully connected to PlanetScale!")
}
