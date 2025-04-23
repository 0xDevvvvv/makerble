package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

var AppConfig *Config

func LoadConfig() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("No .env file found, falling back to system env")
	}

	AppConfig = &Config{
		Port:       getEnv("PORT", "8080"),
		DBHost:     getEnv("DBHOST", "localhost"),
		DBPort:     getEnv("DBPORT", "5432"),
		DBUser:     getEnv("DBUSER", ""),
		DBPassword: getEnv("DBPASSWORD", ""),
		DBName:     getEnv("DBNAME", ""),
		JWTSecret:  getEnv("JWT_SECRET", "secret"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
