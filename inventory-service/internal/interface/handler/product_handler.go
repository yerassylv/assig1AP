package handler

import (
	"context"
	"inventory-service/internal/domain"
	"inventory-service/internal/usecase"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	usecase usecase.ProductUseCase
}

func NewProductHandler(router *gin.Engine, uc usecase.ProductUseCase) {
	handler := &ProductHandler{usecase: uc}

	router.POST("/products", handler.CreateProduct)
	router.GET("/products/:id", handler.GetProductByID)
	router.PATCH("/products/:id", handler.UpdateProduct)
	router.DELETE("/products/:id", handler.DeleteProduct)

}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product domain.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	err := h.usecase.CreateProduct(ctx, &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created"})
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id := c.Param("id")

	product, err := h.usecase.GetProductByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var updatedProduct domain.Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	updatedProduct.ID = id

	err := h.usecase.UpdateProduct(c.Request.Context(), &updatedProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated"})
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	err := h.usecase.DeleteProduct(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
