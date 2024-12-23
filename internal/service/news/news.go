package news

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"stugi/api-gateway/internal/model"
	. "stugi/api-gateway/internal/model"
)

type Service interface {
	GetNews(pageStr, pageSizeStr string) ([]*NewsShortDetailed, error)
	FilterNews(pageStr, pageSizeStr string, filter string) ([]*NewsShortDetailed, error)
	GetNewsDetailed(id string) (*NewsFullDetailed, error)
}

type NewsService struct {
	apiBaseURL string
}

// Конструктор для создания нового экземпляра NewsService
func New(apiBaseURL string) *NewsService {
	return &NewsService{
		apiBaseURL: apiBaseURL,
	}
}

func (s *NewsService) GetNews(pageStr, pageSizeStr string) ([]*NewsShortDetailed, map[string]any) {
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
		return nil, map[string]any{"error": "Unable to fetch news"}
	}
	defer resp.Body.Close()

	var newsResponse struct {
		News       []*NewsShortDetailed `json:"news"`
		Pagination map[string]any       `json:"pagination"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&newsResponse); err != nil {
		return nil, map[string]any{"error": "Unable to decode response"}
	}

	return newsResponse.News, newsResponse.Pagination
}

func (s *NewsService) FilterNews(pageStr, pageSizeStr string, filter string) ([]*model.NewsShortDetailed, map[string]any) {
	// Формируем URL для фильтрации новостей с учетом пагинации
	url := fmt.Sprintf("%s/news/filter?s=%s&page=%s&pageSize=%s", s.apiBaseURL, filter, pageStr, pageSizeStr)

	// Отправляем запрос к внешнему сервису
	resp, err := http.Get(url)
	if err != nil {
		return nil, map[string]any{"error": "Unable to fetch filtered news"}
	}
	defer resp.Body.Close()

	var newsResponse struct {
		News       []*model.NewsShortDetailed `json:"news"`
		Pagination map[string]any             `json:"pagination"`
	}

	// Декодируем ответ
	if err := json.NewDecoder(resp.Body).Decode(&newsResponse); err != nil {
		return nil, map[string]any{"error": "Unable to decode response"}
	}

	// Возвращаем новости с пагинацией
	return newsResponse.News, newsResponse.Pagination
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

	var newsDetails model.NewsFullDetailed
	if err := json.NewDecoder(resp.Body).Decode(&newsDetails); err != nil {
		return nil, fmt.Errorf("unable to decode response: %w", err)
	}

	return &newsDetails, nil
}
