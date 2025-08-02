package config

import (
	"os"
	"strconv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	CORS     CORSConfig
	Env      string
}

type ServerConfig struct {
	Port string
	Host string
}

type DatabaseConfig struct {
	Type     string
	Name     string
	Host     string
	Port     string
	User     string
	Password string
}

type JWTConfig struct {
	Secret string
	Expiry string
}

type CORSConfig struct {
	Origin string
}

var AppConfig *Config

func LoadConfig() {
	AppConfig = &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "8080"),
			Host: getEnv("HOST", "localhost"),
		},
		Database: DatabaseConfig{
			Type:     getEnv("DB_TYPE", "sqlite3"),
			Name:     getEnv("DB_NAME", "ecommerce.db"),
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "3306"),
			User:     getEnv("DB_USER", "root"),
			Password: getEnv("DB_PASSWORD", ""),
		},
		JWT: JWTConfig{
			Secret: getEnv("JWT_SECRET", "your-secret-key-here"),
			Expiry: getEnv("JWT_EXPIRY", "24h"),
		},
		CORS: CORSConfig{
			Origin: getEnv("CORS_ORIGIN", "*"),
		},
		Env: getEnv("ENV", "development"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
} 