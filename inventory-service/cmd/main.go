package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	infraMongo "inventory-service/internal/infrastructure/mongo"
	"inventory-service/internal/interface/handler"
	"inventory-service/internal/usecase"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Mongo connection error:", err)
	}

	db := client.Database("inventorydb")

	productRepo := infraMongo.NewProductRepository(db)
	productUC := usecase.NewProductUseCase(productRepo)

	r := gin.Default()
	handler.NewProductHandler(r, productUC)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
