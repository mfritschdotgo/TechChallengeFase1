package domain

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID        uuid.UUID   `bson:"_id" json:"id"`
	Client    string      `bson:"client" json:"client"`
	Items     []OrderItem `bson:"items" json:"items"`
	Total     float64     `bson:"total" json:"total"`
	Status    int         `bson:"status" json:"status"`
	CreatedAt time.Time   `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time   `bson:"updated_at" json:"updated_at"`
}

type OrderItem struct {
	OrderID   uuid.UUID `bson:"order_id" json:"order_id"`
	ProductID string    `bson:"product_id" json:"product_id"`
	Quantity  int64     `bson:"quantity" json:"quantity"`
	Price     float64   `bson:"price" json:"price"`
}

func NewOrder(client string, items []OrderItem, status int, total float64) (*Order, error) {
	now := time.Now()
	order := &Order{
		ID:        uuid.New(),
		Client:    client,
		Items:     items,
		Total:     total,
		Status:    status,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return order, nil
}
