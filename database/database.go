package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Username string
	Password string
	Name     string
	Host     string
	Port     string
}

func GetConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env")
	}

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	return &Config{
		DB: &DBConfig{
			dbUsername,
			dbPassword,
			dbName,
			dbHost,
			dbPort,
		},
	}
}
