package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"inventory-service/internal/domain"
)

type productRepository struct {
	collection *mongo.Collection
}

// NewProductRepository создает новый репозиторий
func NewProductRepository(db *mongo.Database) domain.ProductRepository {
	return &productRepository{
		collection: db.Collection("products"),
	}
}

// Create — создает новый товар
func (r *productRepository) Create(ctx context.Context, p *domain.Product) error {
	_, err := r.collection.InsertOne(ctx, p)
	return err
}

// GetByID — получить товар по ID
func (r *productRepository) GetByID(ctx context.Context, id string) (*domain.Product, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var product domain.Product
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// Update — обновить товар
func (r *productRepository) Update(ctx context.Context, p *domain.Product) error {
	objID, err := primitive.ObjectIDFromHex(p.ID)
	if err != nil {
		return err
	}

	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{
		"$set": bson.M{
			"name":     p.Name,
			"category": p.Category,
			"price":    p.Price,
			"stock":    p.Stock,
		},
	})
	return err
}

// Delete — удалить товар
func (r *productRepository) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

// List — список всех товаров
func (r *productRepository) List(ctx context.Context, filter map[string]interface{}) ([]*domain.Product, error) {
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []*domain.Product
	for cursor.Next(ctx) {
		var p domain.Product
		if err := cursor.Decode(&p); err != nil {
			return nil, err
		}
		products = append(products, &p)
	}

	return products, nil
}
