package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/mfritschdotgo/techchallenge/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderRepository struct {
	Collection *mongo.Collection
}

func NewOrderRepository(db *mongo.Database) *OrderRepository {
	return &OrderRepository{Collection: db.Collection("orders")}
}

func (pr *OrderRepository) CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	_, err := pr.Collection.InsertOne(ctx, order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (pr *OrderRepository) GetOrders(ctx context.Context, page, limit int) ([]domain.Order, error) {
	var orders []domain.Order
	opts := options.Find().SetSkip(int64((page - 1) * limit)).SetLimit(int64(limit))

	cursor, err := pr.Collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var order domain.Order
		if err := cursor.Decode(&order); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (pr *OrderRepository) GetOrderByID(ctx context.Context, id string) (*domain.Order, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	binaryUUID := primitive.Binary{
		Subtype: 0x00,
		Data:    uuid[:],
	}

	var order domain.Order
	err = pr.Collection.FindOne(ctx, bson.M{"_id": binaryUUID}).Decode(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (pr *OrderRepository) SetStatus(ctx context.Context, id uuid.UUID, status int, description string) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": status, "status_description": description}}
	_, err := pr.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
