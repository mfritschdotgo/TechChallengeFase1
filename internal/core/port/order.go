package port

import (
	"context"

	"github.com/mfritschdotgo/techchallenge/internal/core/domain"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, product *domain.Order) (*domain.Order, error)
	GetOrderByID(ctx context.Context, id string) (*domain.Order, error)
	GetOrders(ctx context.Context, page, pageSize int) ([]domain.Order, error)
}

type OrderService interface {
	CreateOrder(ctx context.Context, product *domain.Order) ([]domain.Order, error)
	GetOrderByID(ctx context.Context, id string) (*domain.Order, error)
	GetOrders(ctx context.Context, page, pageSize int) ([]domain.Order, error)
}
