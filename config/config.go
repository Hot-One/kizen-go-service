package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	ServiceName string
	Environment string
	Host        string
	Port        int

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	EmailHost     string
	EmailPort     int
	EmailUser     string
	EmailPassword string
}

func Load() Config {
	return Config{
		ServiceName: cast.ToString(getOrReturnDefault("SERVICE_NAME", "kizen-go-service")),
		Environment: cast.ToString(getOrReturnDefault("ENVIRONMENT", "development")),
		Host:        cast.ToString(getOrReturnDefault("HOST", "localhost")),
		Port:        cast.ToInt(getOrReturnDefault("PORT", 8080)),

		PostgresHost:     cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost")),
		PostgresPort:     cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432)),
		PostgresUser:     cast.ToString(getOrReturnDefault("POSTGRES_USER", "kizen")),
		PostgresPassword: cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "kizen")),
		PostgresDatabase: cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "kizen")),

		EmailHost:     cast.ToString(getOrReturnDefault("EMAIL_HOST", "smtp.example.com")),
		EmailPort:     cast.ToInt(getOrReturnDefault("EMAIL_PORT", 587)),
		EmailUser:     cast.ToString(getOrReturnDefault("EMAIL_USER", "abulbositkabilov1@gmail.com")),
		EmailPassword: cast.ToString(getOrReturnDefault("EMAIL_PASSWORD", "ouea lhte rhjw spcn")),
	}
}

func getOrReturnDefault(key string, defaultValue any) any {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
