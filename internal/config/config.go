package config

import (
	"os"
)

type Config struct {
	Address     string
	Environment string
	CORSOrigins []string
}

func Load() *Config {
	return &Config{
		Address:     getEnv("API_ADDRESS", ":8080"),
		Environment: getEnv("ENVIRONMENT", "development"),
		CORSOrigins: getCORSOrigins(),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getCORSOrigins() []string {
	origins := os.Getenv("CORS_ORIGINS")
	if origins == "" {
		return []string{"http://localhost:3000"} // Default to UI dev server
	}
	return []string{origins}
}