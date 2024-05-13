package port

import (
	"context"

	"github.com/mfritschdotgo/techchallenge/internal/core/domain"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error)
	GetProductByID(ctx context.Context, id string) (*domain.Product, error)
	GetProducts(ctx context.Context, categoryId string, page, limit int) ([]domain.Product, error)
	UpdateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error)
	DeleteProduct(ctx context.Context, id string) error
}

type ProductService interface {
	CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error)
	GetProductByID(ctx context.Context, id string) (*domain.Product, error)
	GetProducts(ctx context.Context, categoryId string, page, limit int) ([]domain.Product, error)
	UpdateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error)
	DeleteProduct(ctx context.Context, id string) error
}
