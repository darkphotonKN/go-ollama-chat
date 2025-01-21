package main

import (
	"net/http"

	"github.com/darkphotonKN/go-ollama-chat/config"
	"github.com/darkphotonKN/go-ollama-chat/internal/genai"
)

func SetupRoutes(cfg *config.Config) {
	// --- Gen AI ---

	// --- setup ---
	genAIRepo := genai.NewGenAIRepository(cfg)
	genAIService := genai.NewGenAIService(genAIRepo)
	genAIHandler := genai.NewGenAIHandler(genAIService)

	// -- api routes --
	http.HandleFunc("/api/query", genAIHandler.QueryAIHandler)
}
