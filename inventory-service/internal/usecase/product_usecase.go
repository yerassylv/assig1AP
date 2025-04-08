package usecase

import (
	"context"
	"inventory-service/internal/domain"
)

type ProductUseCase interface {
	CreateProduct(ctx context.Context, p *domain.Product) error
	GetProductByID(ctx context.Context, id string) (*domain.Product, error)
	UpdateProduct(ctx context.Context, p *domain.Product) error
	DeleteProduct(ctx context.Context, id string) error
	ListProducts(ctx context.Context, filter map[string]interface{}) ([]*domain.Product, error)
}

type productUseCase struct {
	repo domain.ProductRepository
}

func NewProductUseCase(repo domain.ProductRepository) ProductUseCase {
	return &productUseCase{repo: repo}
}

func (uc *productUseCase) CreateProduct(ctx context.Context, p *domain.Product) error {
	return uc.repo.Create(ctx, p)
}

func (uc *productUseCase) GetProductByID(ctx context.Context, id string) (*domain.Product, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *productUseCase) UpdateProduct(ctx context.Context, p *domain.Product) error {
	return uc.repo.Update(ctx, p)
}

func (uc *productUseCase) DeleteProduct(ctx context.Context, id string) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *productUseCase) ListProducts(ctx context.Context, filter map[string]interface{}) ([]*domain.Product, error) {
	return uc.repo.List(ctx, filter)
}
