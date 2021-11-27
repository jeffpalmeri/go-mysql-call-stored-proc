package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func GetMySQLDB() (db *sql.DB, err error) {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("username")
	password := os.Getenv("password")
	schema := os.Getenv("schema")

	db, err = sql.Open("mysql", username + ":" + password + "@(127.0.0.1:3306)/" + schema + "?parseTime=true")
	return db, err
}
