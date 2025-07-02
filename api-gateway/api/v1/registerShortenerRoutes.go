package v1

import (
	"github.com/Aniketyadav44/s7r/api-gateway/internal/handlers"
	"github.com/Aniketyadav44/s7r/api-gateway/internal/middlewares"
	"github.com/Aniketyadav44/s7r/api-gateway/internal/services"
	pb "github.com/Aniketyadav44/s7r/api-gateway/proto/shortener"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func registerShortenerRoutes(router *gin.Engine, redis *redis.Client, shortenerClient *pb.ShortenerClient) {
	shortenerService := services.NewShortenerService(redis, shortenerClient)
	shortenerHandler := handlers.NewShortenerHandler(shortenerService)

	router.GET("/:short", shortenerHandler.RedirectFromShort)

	v1 := router.Group("/api/v1")
	v1.Use(middlewares.AuthMiddleware())
	{
		v1.POST("/create", shortenerHandler.CreateShortUrl)
		v1.POST("/delete", shortenerHandler.DeleteShortUrl)
	}
}
