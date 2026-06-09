package handlers

import (
	"car-rental-system/backend/internal/models"
	"car-rental-system/backend/internal/services"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CarHandler struct {
	service *services.CarService
}

func NewCarHandler(service *services.CarService) *CarHandler {
	return &CarHandler{
		service: service,
	}
}

func (h *CarHandler) GetAll(c *gin.Context) {

	cars, err := h.service.GetAll()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, cars)
}
func (h *CarHandler) GetByID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid id",
		})
		return
	}

	car, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "car not found",
		})
		return
	}

	c.JSON(200, car)
}
func (h *CarHandler) Create(c *gin.Context) {

	var car models.Car

	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.service.Create(&car)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, car)
}

func (h *CarHandler) Update(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid id",
		})
		return
	}

	var car models.Car

	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	car.ID = id

	err = h.service.Update(&car)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, car)
}
func (h *CarHandler) Delete(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid id",
		})
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "deleted",
	})
}
