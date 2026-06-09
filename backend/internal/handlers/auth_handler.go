package handlers

import (
	"car-rental-system/backend/internal/services"
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

	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.service.Register(
		req.FullName,
		req.Email,
		req.Password,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "user created",
	})
}

func (h *AuthHandler) Login(c *gin.Context) {

	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := h.service.Login(
		req.Email,
		req.Password,
	)

	if err != nil {
		c.JSON(401, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
