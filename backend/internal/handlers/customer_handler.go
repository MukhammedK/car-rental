package handlers

import (
	"car-rental-system/backend/internal/models"
	"car-rental-system/backend/internal/services"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CustomerHandler struct {
	service *services.CustomerService
}

func NewCustomerHandler(
	service *services.CustomerService,
) *CustomerHandler {
	return &CustomerHandler{
		service: service,
	}
}
func (h *CustomerHandler) GetAll(c *gin.Context) {

	customers, err := h.service.GetAll()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, customers)
}
func (h *CustomerHandler) GetByID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid id",
		})
		return
	}

	customer, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "customer not found",
		})
		return
	}

	c.JSON(200, customer)
}
func (h *CustomerHandler) Create(c *gin.Context) {

	var customer models.Customer

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.service.Create(&customer)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, customer)
}
func (h *CustomerHandler) Update(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid id",
		})
		return
	}

	var customer models.Customer

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	customer.ID = id

	err = h.service.Update(&customer)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, customer)
}
func (h *CustomerHandler) Delete(c *gin.Context) {

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
