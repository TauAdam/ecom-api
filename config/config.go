package config

import (
	"fmt"
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
	JWTSecret            string
}

func NewConfig() Config {
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}
	return Config{
		PublicHost:           getEnvVar("PUBLIC_HOST", "http://localhost"),
		Port:                 getEnvVar("PORT", "8080"),
		DBUser:               getEnvVar("DB_USER", "myuser"),
		DBPassword:           getEnvVar("DB_PASSWORD", "mypass"),
		DBName:               getEnvVar("DB_NAME", "mydb"),
		DBAddress:            fmt.Sprintf("%s:%s", getEnvVar("DB_HOST", "localhost"), getEnvVar("DB_PORT", "3306")),
		JWTExpirationSeconds: getEnvAsInteger("JWT_EXPIRATION", 3600*24),
		JWTSecret:            getEnvVar("JWT_SECRET", "super-secret-token"),
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
