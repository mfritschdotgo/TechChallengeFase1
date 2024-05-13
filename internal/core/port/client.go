package port

import (
	"context"

	"github.com/mfritschdotgo/techchallenge/internal/core/domain"
)

type ClientRepository interface {
	CreateClient(ctx context.Context, client *domain.Client) (*domain.Client, error)
	GetClientByCPF(ctx context.Context, cpf string) (*domain.Client, error)
}

type ClientService interface {
	CreateClient(ctx context.Context, client *domain.Client) (*domain.Client, error)
	GetClientByCPF(ctx context.Context, cpf string) (*domain.Client, error)
}
