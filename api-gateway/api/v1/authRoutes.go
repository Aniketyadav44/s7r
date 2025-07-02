package v1

import (
	"database/sql"

	"github.com/Aniketyadav44/s7r/api-gateway/internal/handlers"
	"github.com/Aniketyadav44/s7r/api-gateway/internal/services"
	"github.com/gin-gonic/gin"
)

func registerAuthRoutes(router *gin.Engine, db *sql.DB) {
	authService := services.NewAuthService(db)
	authHandler := handlers.NewAuthHandler(authService)

	v1 := router.Group("/api/v1")
	{
		v1.POST("/register", authHandler.Register)
		v1.POST("/login", authHandler.Login)
	}
}
