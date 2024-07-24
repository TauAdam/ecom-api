package config

import (
	"fmt"
	"os"
)

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBAddress  string
	DBPassword string
	DBName     string
}

func NewConfig() Config {
	return Config{
		PublicHost: getEnvVar("PUBLIC_HOST", "http://localhost"),
		Port:       getEnvVar("PORT", "8080"),
		DBUser:     getEnvVar("DB_USER", "admin"),
		DBPassword: getEnvVar("DB_PASSWORD", "password"),
		DBName:     getEnvVar("DB_NAME", "ecommerce"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnvVar("DB_HOST", "localhost"), getEnvVar("DB_PORT", "3303")),
	}
}

func getEnvVar(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
