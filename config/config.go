package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

// Envs it's one instance that contains all environment variables
var Envs = NewConfig()

type Config struct {
	PublicHost           string
	Port                 string
	DBUser               string
	DBAddress            string
	DBPassword           string
	DBName               string
	JWTExpirationSeconds int64
}

func NewConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return Config{
		PublicHost:           getEnvVar("PUBLIC_HOST", "http://localhost"),
		Port:                 getEnvVar("PORT", "8080"),
		DBUser:               getEnvVar("DB_USER", "admin"),
		DBPassword:           getEnvVar("DB_PASSWORD", "password"),
		DBName:               getEnvVar("DB_NAME", "ecommerce"),
		DBAddress:            fmt.Sprintf("%s:%s", getEnvVar("DB_HOST", "localhost"), getEnvVar("DB_PORT", "3303")),
		JWTExpirationSeconds: getEnvAsInteger("JWT_EXPIRATION", 3600*24),
	}
}

func getEnvVar(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInteger(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}
