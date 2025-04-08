package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order-service/internal/domain"
	"order-service/internal/usecase"
)

type OrderHandler struct {
	usecase usecase.OrderUseCase
}

func NewOrderHandler(router *gin.Engine, uc usecase.OrderUseCase) {
	handler := &OrderHandler{usecase: uc}

	router.POST("/orders", handler.CreateOrder)
	router.GET("/orders/:id", handler.GetOrderByID)
	router.PATCH("/orders/:id", handler.UpdateOrderStatus)
	router.DELETE("/orders/:id", handler.DeleteOrder)
	router.GET("/orders", handler.ListOrders)
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var order domain.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	err := h.usecase.CreateOrder(c.Request.Context(), &order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order created"})
}

func (h *OrderHandler) GetOrderByID(c *gin.Context) {
	id := c.Param("id")

	order, err := h.usecase.GetOrderByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) UpdateOrderStatus(c *gin.Context) {
	id := c.Param("id")
	var updatedOrder domain.Order

	if err := c.ShouldBindJSON(&updatedOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	updatedOrder.ID = id

	err := h.usecase.UpdateOrderStatus(c.Request.Context(), &updatedOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order status updated"})
}

func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")

	err := h.usecase.DeleteOrder(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted"})
}

func (h *OrderHandler) ListOrders(c *gin.Context) {
	orders, err := h.usecase.ListOrders(c.Request.Context(), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}
