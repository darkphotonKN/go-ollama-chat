package genai

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/darkphotonKN/go-ollama-chat/config"
)

type GenAIRepository struct {
	cfg *config.Config
}

func NewGenAIRepository(repo *GenAIRepository) *GenAIService {
	return &GenAIService{
		repo: repo,
	}
}

func (r GenAIRepository) QueryOllama(prompt string) (string, error) {
	payload := ollamaRequest{
		Model:  r.cfg.ModelName,
		Prompt: prompt,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(r.cfg.OllamaURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("non-OK HTTP status: " + resp.Status)
	}

	var result ollamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Response, nil
}
