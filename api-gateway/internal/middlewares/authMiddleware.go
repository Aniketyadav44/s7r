package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Aniketyadav44/s7r/api-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			utils.ErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("please provide Authorization token in headers"))
			c.Abort()
			return
		}

		splittedHeader := strings.Split(authHeader, " ")
		if len(splittedHeader) < 2 || (len(splittedHeader) >= 1 && splittedHeader[0] != "Bearer") {
			utils.ErrorResponse(c, http.StatusUnauthorized, fmt.Errorf("invalid bearer token"))
			c.Abort()
			return
		}

		authToken := splittedHeader[1]
		userId, err := utils.ValidateToken(authToken)
		if err != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, err)
			c.Abort()
			return
		}

		c.Set("user_id", userId)
		c.Next()
	}
}
