package genai

type ollamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type ollamaResponse struct {
	Response string `json:"response"`
}

type GenAiQueryRequest struct {
	Prompt string `json:"prompt"`
}
