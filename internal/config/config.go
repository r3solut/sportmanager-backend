package config

import (
	"os"
	"://github.com"
)

type Config struct {
	DBHost, DBPort, DBUser, DBPassword, DBName, DBSSLMode, JWTSecret, JWTAccessTTL, JWTRefreshTTL, ServerPort string
}

func Load() (*Config, error) {
	_ = godotenv.Load()
	return &Config{
		DBHost: getEnv("DB_HOST", "localhost"), DBPort: getEnv("DB_PORT", "5432"), DBUser: getEnv("DB_USER", "postgres"), DBPassword: getEnv("DB_PASSWORD", "secret"), DBName: getEnv("DB_NAME", "sportmanager"), DBSSLMode: getEnv("DB_SSLMODE", "disable"), JWTSecret: getEnv("JWT_SECRET", "default-secret"), JWTAccessTTL: getEnv("JWT_ACCESS_TTL", "15m"), JWTRefreshTTL: getEnv("JWT_REFRESH_TTL", "720h"), ServerPort: getEnv("SERVER_PORT", "8080"),
	}, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists { return value }
	return fallback
}
