package app

import (
	"stugi/api-gateway/internal/service"

	"github.com/gin-gonic/gin"
)

func New(gateway *service.Gateway) *app {
	return &app{
		gateway: gateway,
	}
}

type app struct {
	gateway *service.Gateway
}

func (a *app) Run() error {
	router := gin.Default()
	// 	запрос получения списка новостей;
	router.GET("/news", getNews)
	// запрос фильтра новостей;
	router.GET("/news/filter", getNews)
	// запрос получения детальной новости;
	router.GET("/news/:id", getNewsByID)
	// запрос добавления комментария.
	router.POST("/comments", postComment)

	err := router.Run(":8080")

	return err
}

func getNews(c *gin.Context) {
	c.JSON(200, gin.H{
		"news": []map[string]any{
			{
				"id":          1,
				"title":       "title1",
				"description": "description1",
			},
			{
				"id":          2,
				"title":       "title2",
				"description": "description2",
			},
		},
	})
}

func getNewsByID(c *gin.Context) {
	c.JSON(200, gin.H{
		"news": map[string]any{
			"id":          1,
			"title":       "title1",
			"description": "description1",
		},
		"comments": []map[string]any{
			{
				"id":   1,
				"text": "comments 1",
			},
			{
				"id":   2,
				"text": "comments 2",
			},
		},
	})
}

func postComment(c *gin.Context) {
	c.JSON(200, gin.H{
		"news": map[string]any{
			"id":          1,
			"title":       "title1",
			"description": "description1",
		},
		"comments": []map[string]any{
			{
				"id":   1,
				"text": "comments 1",
			},
			{
				"id":   2,
				"text": "comments 2",
			},
		},
	})
}
