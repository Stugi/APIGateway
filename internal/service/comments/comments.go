package comments

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"stugi/api-gateway/internal/model"
	. "stugi/api-gateway/internal/model"
)

type Service interface {
	AddComment(comment *Comment) error
}

type CommentsService struct {
	apiBaseURL string
}

func New(apiBaseURL string) *CommentsService {
	return &CommentsService{
		apiBaseURL: apiBaseURL,
	}
}

// Метод для добавления комментария
func (s *CommentsService) AddComment(comment *model.Comment) error {
	// Формируем URL для добавления комментария
	url := fmt.Sprintf("%s/comments", s.apiBaseURL)

	// Преобразуем комментарий в JSON
	commentJSON, err := json.Marshal(comment)
	if err != nil {
		return fmt.Errorf("unable to marshal comment: %w", err)
	}

	// Отправляем POST-запрос к внешнему сервису
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(commentJSON))
	if err != nil {
		return fmt.Errorf("unable to add comment: %w", err)
	}
	defer resp.Body.Close()

	// Проверяем статус код ответа
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to add comment, status code: %d", resp.StatusCode)
	}

	// Успешное добавление комментария
	return nil
}
