package handlers

import (
	"fmt"
	"net/http"

	"github.com/Aniketyadav44/s7r/api-gateway/internal/services"
	"github.com/Aniketyadav44/s7r/api-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

type GetterHandler struct {
	service *services.GetterService
}

func NewGetterHandler(service *services.GetterService) *GetterHandler {
	return &GetterHandler{
		service: service,
	}
}

func (h *GetterHandler) GetUserURLs(c *gin.Context) {
	userId := c.GetString("user_id")
	if userId == "" {
		utils.ErrorResponse(c, -1, fmt.Errorf("cannot retrieve userId"))
		return
	}
	urls, err := h.service.GetAllUrls(userId)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success", "urls": urls})
}

func (h *GetterHandler) GetURLClicks(c *gin.Context) {
	type ClicksRequest struct {
		ShortCode string `json:"short_code" binding:"required"`
	}

	var reqBody ClicksRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	count, err := h.service.GetClicks(reqBody.ShortCode)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success", "clicks": count})
}

func (h *GetterHandler) GetURLHistory(c *gin.Context) {
	type ClicksRequest struct {
		ShortCode string `json:"short_code" binding:"required"`
	}

	var reqBody ClicksRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	history, err := h.service.GetHistory(reqBody.ShortCode)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success", "history": history})
}
