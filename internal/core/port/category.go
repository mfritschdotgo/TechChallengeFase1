package port

import (
	"context"

	"github.com/mfritschdotgo/techchallenge/internal/core/domain"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error)
	GetCategoryByID(ctx context.Context, id string) (*domain.Category, error)
	GetCategories(ctx context.Context, page, limit int) ([]domain.Category, error)
	UpdateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error)
	DeleteCategory(ctx context.Context, id string) error
	PatchCategory(ctx context.Context, category *domain.Category) (*domain.Category, error)
}

type CategoryService interface {
	CreateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error)
	GetCategoryByID(ctx context.Context, id string) (*domain.Category, error)
	GetCategories(ctx context.Context, page, limit int) ([]domain.Category, error)
	UpdateCategory(ctx context.Context, category *domain.Category) (*domain.Category, error)
	DeleteCategory(ctx context.Context, id string) error
	PatchCategory(ctx context.Context, category *domain.Category) (*domain.Category, error)
}
