package config

import (
	"os"
)

type Config struct {
	Port    string
	LogPath string
	Env     string
}

func Load() *Config {
	return &Config{
		Port:    getEnv("PORT", "8080"),
		LogPath: getEnv("LOG_PATH", "server.log"),
		Env:     getEnv("ENV", "dev"),
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
