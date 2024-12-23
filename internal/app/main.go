package app

import (
	"strconv"
	"stugi/api-gateway/internal/app/middleware"
	"stugi/api-gateway/internal/model"
	srvComments "stugi/api-gateway/internal/service/comments"
	srvNews "stugi/api-gateway/internal/service/news"

	"github.com/gin-gonic/gin"
)

func New(service srvNews.Service, serviceComments srvComments.Service) *app {
	return &app{
		serviceNews:     service,
		serviceComments: serviceComments,
	}
}

type app struct {
	serviceNews     srvNews.Service
	serviceComments srvComments.Service
}

func (a *app) Run() error {
	router := gin.Default()
	// Добавляем middleware для поддержки request_id
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.LoggerMiddleware())
	// 	запрос получения списка новостей;
	router.GET("/news", a.getNews)
	// запрос получения детальной новости;
	router.GET("/news/:id", a.getNewsByID)
	// запрос добавления комментария.
	router.POST("/comments", a.postComment)

	err := router.Run(":8080")

	return err
}

func (a *app) getNews(c *gin.Context) {
	// Получаем параметры запроса для поиска и пагинации
	search := c.DefaultQuery("s", "")
	pageStr := c.DefaultQuery("page", "1")          // Если нет параметра "page", по умолчанию 1
	pageSizeStr := c.DefaultQuery("pageSize", "10") // Если нет параметра "pageSize", по умолчанию 10

	// Используем сервис для получения новостей
	news, pagination := getPaginatedNews(a.serviceNews, search, pageStr, pageSizeStr)

	// Возвращаем новость с пагинацией
	c.JSON(200, gin.H{
		"news":       news,
		"pagination": pagination,
	})
}

func (a *app) getNewsByID(c *gin.Context) {
	id := c.Param("id")
	detailedNews, err := a.serviceNews.GetNewsDetailed(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Unable to fetch news details"})
		return
	}

	c.JSON(200, gin.H{
		"news": detailedNews,
	})
}

func (a *app) postComment(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Comment added successfully",
	})
}

// Дополнительная функция для пагинации
func getPaginatedNews(service srvNews.Service, search, pageStr, pageSizeStr string) ([]*model.NewsShortDetailed, map[string]any) {
	news, err := service.FilterNews(pageStr, pageSizeStr, search)
	if err != nil {
		return nil, nil
	}

	// Преобразуем параметры в целые числа
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	// Реализуем пагинацию
	start := (page - 1) * pageSize
	end := start + pageSize

	// Если запрашиваемая страница выходит за пределы, ограничиваем её
	if start > len(news) {
		return []*model.NewsShortDetailed{}, map[string]any{
			"page":     page,
			"pageSize": pageSize,
			"total":    len(news),
		}
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
