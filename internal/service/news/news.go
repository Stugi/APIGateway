package news

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"stugi/api-gateway/internal/model"
	. "stugi/api-gateway/internal/model"
	"stugi/api-gateway/internal/service/comments"
)

type Service interface {
	GetNews(pageStr, pageSizeStr string) ([]*NewsShortDetailed, error)
	FilterNews(pageStr, pageSizeStr, filter string) ([]*NewsShortDetailed, error)
	GetNewsDetailed(id string) (*NewsFullDetailed, error)
}

type NewsService struct {
	apiBaseURL      string
	serviceComments *comments.CommentsService
}

// Конструктор для создания нового экземпляра NewsService
func New(apiBaseURL string, serviceComments *comments.CommentsService) *NewsService {
	return &NewsService{
		apiBaseURL:      apiBaseURL,
		serviceComments: serviceComments,
	}
}

func (s *NewsService) GetNews(pageStr, pageSizeStr string) ([]*NewsShortDetailed, error) {
	// Формируем URL для запроса к внешнему сервису
	url := fmt.Sprintf("%s/news", s.apiBaseURL)

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	url += fmt.Sprintf("&page=%d&pageSize=%d", page, pageSize)

	// Отправляем запрос к внешнему сервису
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var news []*NewsShortDetailed

	if err := json.NewDecoder(resp.Body).Decode(&news); err != nil {
		return nil, err
	}

	return news, nil
}

func (s *NewsService) FilterNews(pageStr, pageSizeStr, filter string) ([]*model.NewsShortDetailed, error) {
	// Формируем URL для фильтрации новостей с учетом пагинации
	url := fmt.Sprintf("%s/news/filter?s=%s&page=%s&pageSize=%s", s.apiBaseURL, filter, pageStr, pageSizeStr)

	// Отправляем запрос к внешнему сервису
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var news []*NewsShortDetailed

	// Декодируем ответ
	if err := json.NewDecoder(resp.Body).Decode(&news); err != nil {
		return nil, err
	}

	return news, nil
}

func (s *NewsService) GetNewsDetailed(id string) (*NewsFullDetailed, error) {
	// Формируем URL для получения подробной новости
	url := fmt.Sprintf("%s/news/%s", s.apiBaseURL, id)

	// Отправляем запрос к внешнему сервису
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch news details: %w", err)
	}
	defer resp.Body.Close()

	var newsDetails NewsFullDetailed
	if err := json.NewDecoder(resp.Body).Decode(&newsDetails); err != nil {
		return nil, fmt.Errorf("unable to decode response: %w", err)
	}

	// Добавляем комментарии к новости
	comments, err := s.serviceComments.GetCommentsByNewsID(id)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch comments: %w", err)
	}
	newsDetails.Comments = &comments

	return &newsDetails, nil
}
