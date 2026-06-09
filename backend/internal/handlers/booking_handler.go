package handlers

import (
	"car-rental-system/backend/internal/models"
	"car-rental-system/backend/internal/services"
	"github.com/gin-gonic/gin"
	"strconv"
)

type BookingHandler struct {
	service *services.BookingService
}

type UpdateBookingStatusRequest struct {
	Status string `json:"status"`
}

func NewBookingHandler(
	service *services.BookingService,
) *BookingHandler {
	return &BookingHandler{
		service: service,
	}
}

func (h *BookingHandler) Create(c *gin.Context) {

	var booking models.Booking

	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	booking.Status = "new"
	booking.Source = "crm"

	err := h.service.Create(&booking)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, booking)
}

func (h *BookingHandler) GetAll(c *gin.Context) {

	bookings, err := h.service.GetAll()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, bookings)
}

func (h *BookingHandler) GetByID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid id",
		})
		return
	}

	booking, err := h.service.GetByID(id)

	if err != nil {
		c.JSON(404, gin.H{
			"error": "booking not found",
		})
		return
	}

	c.JSON(200, booking)
}

func (h *BookingHandler) UpdateStatus(
	c *gin.Context,
) {

	id, err := strconv.Atoi(
		c.Param("id"),
	)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid id",
		})
		return
	}

	var req UpdateBookingStatusRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.service.UpdateStatus(
		id,
		req.Status,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "status updated",
	})
}
