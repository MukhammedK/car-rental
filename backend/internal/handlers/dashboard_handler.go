package handlers

import (
	"car-rental-system/backend/internal/services"
	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	service *services.DashboardService
}

func NewDashboardHandler(
	service *services.DashboardService,
) *DashboardHandler {
	return &DashboardHandler{
		service: service,
	}
}
func (h *DashboardHandler) GetStats(
	c *gin.Context,
) {

	stats, err := h.service.GetStats()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, stats)
}
