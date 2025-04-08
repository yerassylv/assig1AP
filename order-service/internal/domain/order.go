package domain

import "time"
import "context"

type Order struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	UserID    string    `json:"user_id" bson:"user_id"`
	Products  []Product `json:"products" bson:"products"`
	Status    string    `json:"status" bson:"status"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

type Product struct {
	ID       string  `json:"id" bson:"id"`
	Name     string  `json:"name" bson:"name"`
	Price    float64 `json:"price" bson:"price"`
	Quantity int     `json:"quantity" bson:"quantity"`
}

type OrderRepository interface {
	Create(ctx context.Context, o *Order) error
	GetByID(ctx context.Context, id string) (*Order, error)
	Update(ctx context.Context, o *Order) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filter map[string]interface{}) ([]*Order, error)
}
