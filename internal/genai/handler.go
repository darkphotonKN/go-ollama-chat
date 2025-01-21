package genai

import (
	"encoding/json"
	"log"
	"net/http"
)

type GenAIHandler struct {
	service *GenAIService
}

func NewGenAIHandler(service *GenAIService) *GenAIHandler {
	return &GenAIHandler{
		service: service,
	}
}

func (h *GenAIHandler) QueryAIHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		Prompt string `json:"prompt"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	response, err := h.service.QueryAIService(request.Prompt)
	if err != nil {
		log.Printf("Error querying AI: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"response": response,
	})
}
