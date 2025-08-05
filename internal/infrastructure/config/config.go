package config

import (
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
}

func Load() (*Config, error) {
	godotenv.Load()

	config := &Config{
		Port:       getEnv("PORT", "8080"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "admin"),
		DBPassword: getEnv("DB_PASSWORD", "admin"),
		DBName:     getEnv("DB_NAME", "products"),
	}

	// config := &Config{
	// 	Port:       getEnv("PORT", "8080"),
	// 	DBHost:     getEnv("DB_HOST", "db.ihjcvjxthqzcoukulojk.supabase.co"),
	// 	DBPort:     getEnv("DB_PORT", "5432"),
	// 	DBUser:     getEnv("DB_USER", "postgres"),
	// 	DBPassword: getEnv("DB_PASSWORD", "admin@123"),
	// 	DBName:     getEnv("DB_NAME", "postgres"),
	// }

	return config, nil
}

func getEnv(key, defaultValue string) string {

	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
