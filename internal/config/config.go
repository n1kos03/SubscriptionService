package config

import "os"

type DBConfig struct {
	DBPort string
	DBUser string
	DBHost string
	DBSSLMode string
	DBPassword string
	DBName string
}

func Load() *DBConfig {
	return &DBConfig{
		DBPort: getEnv("POSTGRES_PORT", "5431"),
		DBUser: getEnv("POSTGRES_USER", "user"),
		DBHost: getEnv("POSTGRES_HOST", "localhost"),
		DBSSLMode: getEnv("POSTGRES_SSLMODE", "disable"),
		DBPassword: getEnv("POSTGRES_PASSWORD", "123"),
		DBName: getEnv("POSTGRES_DB", "postgres"),
	}
}

func getEnv(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return defaultValue
}