package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Host string
	Port int

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
}

func Load() *Config {
	return &Config{
		Host: cast.ToString(getOrReturnDefault("HOST", "localhost")),
		Port: cast.ToInt(getOrReturnDefault("PORT", 8080)),

		PostgresHost:     cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost")),
		PostgresPort:     cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432)),
		PostgresUser:     cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres")),
		PostgresPassword: cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "postgres")),
		PostgresDatabase: cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "kizen")),
	}
}

func getOrReturnDefault(key string, defaultValue any) any {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
