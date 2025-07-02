package v1

import (
	"github.com/Aniketyadav44/s7r/api-gateway/internal/handlers"
	"github.com/Aniketyadav44/s7r/api-gateway/internal/middlewares"
	"github.com/Aniketyadav44/s7r/api-gateway/internal/services"
	pb "github.com/Aniketyadav44/s7r/api-gateway/proto/getter"
	"github.com/gin-gonic/gin"
)

func registerGetterRoutes(router *gin.Engine, getterClient *pb.GetterClient) {
	getterService := services.NewGetterService(getterClient)
	getterHandler := handlers.NewGetterHandler(getterService)

	v1 := router.Group("/api/v1/url")
	v1.Use(middlewares.AuthMiddleware())
	{
		v1.GET("/all", getterHandler.GetUserURLs)
		v1.POST("/clicks", getterHandler.GetURLClicks)
		v1.POST("/history", getterHandler.GetURLHistory)
	}
}
