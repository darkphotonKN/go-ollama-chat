package config

import (
	"os"
)

type Config struct {
	OllamaURL string
	ModelName string
	Port      string
}

func LoadConfig() Config {
	return Config{
		OllamaURL: getEnv("OLLAMA_URL", "http://localhost:11434/api/chat"),
		ModelName: getEnv("OLLAMA_MODEL", "llama2"),
		Port:      getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
