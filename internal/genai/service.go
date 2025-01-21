package genai

type GenAIService struct {
	repo *GenAIRepository
}

func NewGenAIService(repo *GenAIRepository) *GenAIService {
	return &GenAIService{
		repo: repo,
	}
}

func (s *GenAIService) QueryAIService(prompt string) (string, error) {
	return s.repo.QueryOllama(prompt)
}
