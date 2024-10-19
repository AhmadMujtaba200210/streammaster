package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path/filepath"
	"streammaster/config"
)

func main() {
	env := os.Getenv("ENV_APP")
	if env == "" {
		env = "dev"
	}
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = "debug"
	}
	gin.SetMode(ginMode)
	configFile := filepath.Join("config", "config-"+env+".yml")
	cfg, err := config.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})
	if err := router.Run(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
