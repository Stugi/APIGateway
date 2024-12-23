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
	GetCommentsByNewsID(newsID string) ([]Comment, error)
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

// Метод для получения комментариев к новости
func (s *CommentsService) GetCommentsByNewsID(newsID string) ([]Comment, error) {
	// Формируем URL для получения комментариев
	url := fmt.Sprintf("%s/comments/%s", s.apiBaseURL, newsID)

	// Отправляем GET-запрос к внешнему сервису
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch comments: %w", err)
	}
	defer resp.Body.Close()

	// Проверяем статус код ответа
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch comments, status code: %d", resp.StatusCode)
	}

	// Декодируем ответ
	var comments []Comment
	if err := json.NewDecoder(resp.Body).Decode(&comments); err != nil {
		return nil, fmt.Errorf("unable to decode response: %w", err)
	}

	return comments, nil
}
