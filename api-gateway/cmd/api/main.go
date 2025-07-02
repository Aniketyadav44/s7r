package main

import (
	"log"

	v1 "github.com/Aniketyadav44/s7r/api-gateway/api/v1"
	"github.com/Aniketyadav44/s7r/api-gateway/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("error in loading config: ", err.Error())
	}

	defer cfg.Db.Close()
	defer cfg.Redis.Close()

	router := gin.Default()

	v1.RegisterRoutes(router, cfg)

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal("error in starting server: ", err.Error())
	}
}
