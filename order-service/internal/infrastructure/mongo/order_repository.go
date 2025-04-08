package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"order-service/internal/domain"
)

type orderRepository struct {
	collection *mongo.Collection
}

// NewOrderRepository создает новый репозиторий для заказов
func NewOrderRepository(db *mongo.Database) domain.OrderRepository {
	return &orderRepository{
		collection: db.Collection("orders"),
	}
}

// Create — создает новый заказ
func (r *orderRepository) Create(ctx context.Context, o *domain.Order) error {
	_, err := r.collection.InsertOne(ctx, o)
	return err
}

// GetByID — получает заказ по ID
func (r *orderRepository) GetByID(ctx context.Context, id string) (*domain.Order, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var order domain.Order
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// Update — обновляет статус заказа
func (r *orderRepository) Update(ctx context.Context, o *domain.Order) error {
	objID, err := primitive.ObjectIDFromHex(o.ID)
	if err != nil {
		return err
	}

	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{
		"$set": bson.M{
			"status": o.Status,
		},
	})
	return err
}

// Delete — удаляет заказ по ID
func (r *orderRepository) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

// List — получает список заказов
func (r *orderRepository) List(ctx context.Context, filter map[string]interface{}) ([]*domain.Order, error) {
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []*domain.Order
	for cursor.Next(ctx) {
		var o domain.Order
		if err := cursor.Decode(&o); err != nil {
			return nil, err
		}
		orders = append(orders, &o)
	}

	return orders, nil
}
