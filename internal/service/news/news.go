package news

import (
	"strconv"
	"strings"
	. "stugi/api-gateway/internal/model"
)

type Service interface {
	GetNews() ([]*NewsShortDetailed, error)
	FilterNews(filter string) ([]*NewsShortDetailed, error)
	GetNewsDetailed(id string) (*NewsFullDetailed, error)
}

type NewsService struct{}

func (s *NewsService) GetNews(search, pageStr, pageSizeStr string) ([]NewsShortDetailed, map[string]any) {
	// Преобразуем параметры в целые числа
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	// Пример данных (обычно данные берутся из базы данных или другого сервиса)
	news := []NewsShortDetailed{
		{ID: 1, Title: "title1", Description: "description1"},
		{ID: 2, Title: "title2", Description: "description2"},
		{ID: 3, Title: "special title3", Description: "description3"},
		{ID: 4, Title: "title4", Description: "description4"},
		{ID: 5, Title: "title5", Description: "description5"},
	}

	// Если параметр поиска есть, фильтруем новости по названию
	if search != "" {
		var filteredNews []NewsShortDetailed
		for _, n := range news {
			if strings.Contains(strings.ToLower(n.Title), strings.ToLower(search)) {
				filteredNews = append(filteredNews, n)
			}
		}
		news = filteredNews
	}

	// Реализуем пагинацию
	start := (page - 1) * pageSize
	end := start + pageSize

	// Если запрашиваемая страница выходит за пределы, ограничиваем её
	if start > len(news) {
		return []NewsShortDetailed{}, map[string]any{"page": page, "pageSize": pageSize, "total": len(news)}
	}

	if end > len(news) {
		end = len(news)
	}

	// Пагинация
	pagination := map[string]any{
		"page":     page,
		"pageSize": pageSize,
		"total":    len(news),
	}

	return news[start:end], pagination
}

func (s *NewsService) FilterNews(filter string) ([]*NewsShortDetailed, error) {
	// Пример данных (обычно данные берутся из базы данных или другого сервиса)
	news := []*NewsShortDetailed{
		{ID: 1, Title: "title1", Description: "description1"},
		{ID: 2, Title: "title2", Description: "description2"},
		{ID: 3, Title: "special title3", Description: "description3"},
		{ID: 4, Title: "title4", Description: "description4"},
		{ID: 5, Title: "title5", Description: "description5"},
	}

	// Фильтруем новости по заголовку
	var filteredNews []*NewsShortDetailed
	for _, n := range news {
		if strings.Contains(strings.ToLower(n.Title), strings.ToLower(filter)) {
			filteredNews = append(filteredNews, n)
		}
	}

	return filteredNews, nil
}

func (s *NewsService) GetNewsDetailed(id string) (*NewsFullDetailed, error) {
	// Пример данных (обычно данные берутся из базы данных или другого сервиса)
	detailedNews := &NewsFullDetailed{
		ID:          1,
		Title:       "title1",
		Description: "description1",
		Comments: &[]Comment{
			{ID: 1, Text: "comments 1"},
			{ID: 2, Text: "comments 2"},
		},
	}

	return detailedNews, nil
}
