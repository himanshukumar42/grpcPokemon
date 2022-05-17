package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetConnectionString() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, dbname)

	return connectionString
}
