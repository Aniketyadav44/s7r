package handlers

import (
	"net/http"

	"github.com/Aniketyadav44/s7r/api-gateway/internal/services"
	"github.com/Aniketyadav44/s7r/api-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	type RegisterRequest struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}

	var reqBody RegisterRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	createdUser, err := h.service.Register(reqBody.Name, reqBody.Email, reqBody.Password)
	if err != nil {
		utils.ErrorResponse(c, -1, err)
		return
	}

	jwt, err := utils.GenerateToken(createdUser.Id)
	if err != nil {
		utils.ErrorResponse(c, -1, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "token": jwt, "user": createdUser})
}

func (h *AuthHandler) Login(c *gin.Context) {
	type LoginRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}

	var reqBody LoginRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	user, err := h.service.Login(reqBody.Email, reqBody.Password)
	if err != nil {
		utils.ErrorResponse(c, -1, err)
		return
	}

	jwt, err := utils.GenerateToken(user.Id)
	if err != nil {
		utils.ErrorResponse(c, -1, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged in successfully", "token": jwt, "user": user})
}
