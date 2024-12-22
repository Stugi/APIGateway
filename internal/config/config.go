package config

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

type Config struct {
	Mode string
}

func init() {
	// loads values from .env into the system
	err := gotenv.Load()
	if err != nil {
		log.Println("Error loading .env file", err)
	}
}

func New() *Config {
	return &Config{
		Mode: getEnv("GIN_MODE", gin.DebugMode),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
