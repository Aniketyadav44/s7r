package v1

import (
	"github.com/Aniketyadav44/s7r/api-gateway/internal/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, cfg *config.Config) {
	registerShortenerRoutes(router, cfg.Redis, cfg.ShortenerClient)
	registerGetterRoutes(router, cfg.GetterClient)
	registerAuthRoutes(router, cfg.Db)
}
