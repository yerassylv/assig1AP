package domain

import "context"

type Product struct {
	ID       string  `json:"id" bson:"_id,omitempty"`
	Name     string  `json:"name" bson:"name"`
	Category string  `json:"category" bson:"category"`
	Price    float64 `json:"price" bson:"price"`
	Stock    int     `json:"stock" bson:"stock"`
}

type ProductRepository interface {
	Create(ctx context.Context, p *Product) error
	GetByID(ctx context.Context, id string) (*Product, error)
	Update(ctx context.Context, p *Product) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filter map[string]interface{}) ([]*Product, error)
}
