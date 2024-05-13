package repository

import (
	"context"

	"github.com/mfritschdotgo/techchallenge/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClientRepository struct {
	Collection *mongo.Collection
}

func NewClientRepository(db *mongo.Database) *ClientRepository {
	return &ClientRepository{
		Collection: db.Collection("clients"),
	}
}

func (r *ClientRepository) CreateClient(ctx context.Context, client *domain.Client) (*domain.Client, error) {
	_, err := r.Collection.InsertOne(ctx, client)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (r *ClientRepository) GetClientByCPF(ctx context.Context, cpf string) (*domain.Client, error) {
	var client domain.Client
	err := r.Collection.FindOne(ctx, bson.M{"cpf": cpf}).Decode(&client)
	if err != nil {
		return nil, err
	}
	return &client, nil
}
