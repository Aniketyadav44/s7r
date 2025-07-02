package handlers

import (
	"fmt"
	"net/http"

	"github.com/Aniketyadav44/s7r/api-gateway/internal/services"
	"github.com/Aniketyadav44/s7r/api-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

type ShortenerHandler struct {
	service *services.ShortenerService
}

func NewShortenerHandler(service *services.ShortenerService) *ShortenerHandler {
	return &ShortenerHandler{
		service: service,
	}
}

func (h *ShortenerHandler) RedirectFromShort(c *gin.Context) {
	shortCode, ok := c.Params.Get("short")
	if !ok {
		c.String(http.StatusBadRequest, "invalid short url")
	}
	fmt.Println("calling register")
	url, err := h.service.GetUrlFromShort(shortCode)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println("calling register")
	go h.service.RegisterUrlClick(shortCode)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *ShortenerHandler) CreateShortUrl(c *gin.Context) {
	userId := c.GetString("user_id")
	if userId == "" {
		utils.ErrorResponse(c, -1, fmt.Errorf("cannot retrieve userId"))
		return
	}
	type CreateRequest struct {
		Url string `json:"url" binding:"required"`
	}

	var reqBody CreateRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	shortCode, err := h.service.CreateShortCode(reqBody.Url, userId)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "URL Created successfully",
		"url":        reqBody.Url,
		"short_code": shortCode,
	})
}

func (h *ShortenerHandler) DeleteShortUrl(c *gin.Context) {
	type ClicksRequest struct {
		ShortCode string `json:"short_code" binding:"required"`
	}

	var reqBody ClicksRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	err := h.service.DeleteUrl(reqBody.ShortCode)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "URL deleted successfully", "short_code": reqBody.ShortCode})
}
