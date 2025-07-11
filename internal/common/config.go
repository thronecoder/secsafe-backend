package common

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	RedisAddr  string
	AppPort    string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env")
	}

	return Config{
		DBHost:     getEnv("DB_HOST"),
		DBPort:     getEnv("DB_PORT"),
		DBUser:     getEnv("DB_USER"),
		DBPassword: getEnv("DB_PASSWORD"),
		DBName:     getEnv("DB_NAME"),
		RedisAddr:  getEnv("REDIS_ADDR"),
		AppPort:    getEnv("APP_PORT"),
	}
}

func getEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Missing required env: %s", key)
	}
	return val
}
