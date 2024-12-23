package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Middleware для поддержки сквозного идентификатора
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Пытаемся извлечь X-Request-ID из заголовков запроса
		requestID := c.GetHeader("X-Request-ID")

		// Если заголовка нет, генерируем новый request_id
		if requestID == "" {
			requestID = uuid.New().String() // Генерация нового UUID
		}

		// Добавляем request_id в контекст для дальнейшего использования
		c.Set("request_id", requestID)

		// Логируем request_id
		log.Printf("Request ID: %s", requestID)

		// Передаем управление дальше
		c.Next()
	}
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Логируем начало запроса
		log.Printf("Request started: %s %s", c.Request.Method, c.Request.URL.Path)

		// Передаем управление дальше
		c.Next()

		// Логируем окончание запроса
		log.Printf("Request completed: %s %s", c.Request.Method, c.Request.URL.Path)
	}
}
