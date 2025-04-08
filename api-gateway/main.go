package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	r.Use(gin.Logger())

	r.POST("/orders", proxyToOrderService)
	r.GET("/orders/:id", proxyToOrderService)
	r.PATCH("/orders/:id", proxyToOrderService)
	r.DELETE("/orders/:id", proxyToOrderService)

	r.POST("/products", proxyToInventoryService)
	r.GET("/products/:id", proxyToInventoryService)
	r.PATCH("/products/:id", proxyToInventoryService)
	r.DELETE("/products/:id", proxyToInventoryService)
	r.GET("/products", proxyToInventoryService)

	r.Run(":8082")
}

func proxyToOrderService(c *gin.Context) {
	url := "http://localhost:8081" + c.Request.URL.Path
	req, err := http.NewRequest(c.Request.Method, url, c.Request.Body)
	if err != nil {
		log.Println("Error in creating request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in request forwarding"})
		return
	}

	req.Header = c.Request.Header

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in proxying request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Order Service unavailable"})
		return
	}
	defer resp.Body.Close()

	c.JSON(resp.StatusCode, gin.H{"message": "Request forwarded to Order Service"})
}

func proxyToInventoryService(c *gin.Context) {

	url := "http://localhost:8080" + c.Request.URL.Path

	req, err := http.NewRequest(c.Request.Method, url, c.Request.Body)
	if err != nil {
		log.Println("Error in creating request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in request forwarding"})
		return
	}

	req.Header = c.Request.Header

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in proxying request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Inventory Service unavailable"})
		return
	}
	defer resp.Body.Close()

	c.JSON(resp.StatusCode, gin.H{"message": "Request forwarded to Inventory Service"})
}
