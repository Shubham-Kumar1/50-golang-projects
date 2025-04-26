package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port           string
	WebSocketPath  string
	MaxConnections int
}

func LoadConfig() *Config {
	return &Config{
		Port:           getEnv("PORT", "9090"),
		WebSocketPath:  getEnv("WS_PATH", "/ws"),
		MaxConnections: getEnvAsInt("MAX_CONNECTIONS", 100),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
