package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectDatabase() {
	var err error

	// Charger les variables d'environnement
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Lire DSN depuis les variables d'environnement
	dsn := os.Getenv("DSN")
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}
}
