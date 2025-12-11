package config

import "os"

type Config struct {
	ServerPort     string
	ExternalAPIURL string
}

func Load() *Config {
	return &Config{
		ServerPort:     getEnv("SERVER_PORT", "8080"),
		ExternalAPIURL: getEnv("EXTERNAL_API_URL", "https://jsonplaceholder.typicode.com"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

