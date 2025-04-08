package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"

	orderMongo "order-service/internal/infrastructure/mongo" // используй новое имя для импорта
	"order-service/internal/interface/handler"
	"order-service/internal/usecase"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Mongo connection error:", err)
	}

	db := client.Database("ordersdb")

	orderRepo := orderMongo.NewOrderRepository(db) // используем правильный импорт
	orderUC := usecase.NewOrderUseCase(orderRepo)

	r := gin.Default()
	handler.NewOrderHandler(r, orderUC)

	if err := r.Run(":8081"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
