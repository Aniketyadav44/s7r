package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, code int, err error) {
	if code == -1 {
		code = http.StatusInternalServerError
	}
	c.JSON(code, gin.H{
		"error": err.Error(),
	})
}
