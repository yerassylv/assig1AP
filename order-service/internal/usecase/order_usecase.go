package usecase

import (
	"context"
	"order-service/internal/domain"
)

type OrderUseCase interface {
	CreateOrder(ctx context.Context, o *domain.Order) error
	GetOrderByID(ctx context.Context, id string) (*domain.Order, error)
	UpdateOrderStatus(ctx context.Context, o *domain.Order) error
	DeleteOrder(ctx context.Context, id string) error
	ListOrders(ctx context.Context, filter map[string]interface{}) ([]*domain.Order, error)
}

type orderUseCase struct {
	repo domain.OrderRepository
}

// NewOrderUseCase создаёт новый usecase для заказов
func NewOrderUseCase(repo domain.OrderRepository) OrderUseCase {
	return &orderUseCase{repo: repo}
}

func (uc *orderUseCase) CreateOrder(ctx context.Context, o *domain.Order) error {
	return uc.repo.Create(ctx, o)
}

func (uc *orderUseCase) GetOrderByID(ctx context.Context, id string) (*domain.Order, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *orderUseCase) UpdateOrderStatus(ctx context.Context, o *domain.Order) error {
	return uc.repo.Update(ctx, o)
}

func (uc *orderUseCase) DeleteOrder(ctx context.Context, id string) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *orderUseCase) ListOrders(ctx context.Context, filter map[string]interface{}) ([]*domain.Order, error) {
	return uc.repo.List(ctx, filter)
}
