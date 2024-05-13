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

type ProductRepository struct {
	Collection *mongo.Collection
}

func NewProductRepository(db *mongo.Database) *ProductRepository {
	return &ProductRepository{Collection: db.Collection("products")}
}

func (pr *ProductRepository) CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	_, err := pr.Collection.InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (pr *ProductRepository) GetProductByID(ctx context.Context, id string) (*domain.Product, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	binaryUUID := primitive.Binary{
		Subtype: 0x00,
		Data:    uuid[:],
	}

	var product domain.Product
	err = pr.Collection.FindOne(ctx, bson.M{"_id": binaryUUID}).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (pr *ProductRepository) UpdateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	uuid, err := uuid.Parse(product.ID.String())
	if err != nil {
		return nil, err
	}

	binaryUUID := primitive.Binary{
		Subtype: 0x00,
		Data:    uuid[:],
	}

	filter := bson.M{"_id": binaryUUID}
	update := bson.M{"$set": product}
	_, err = pr.Collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return nil, err
	}
	return product, nil
}

func (pr *ProductRepository) DeleteProduct(ctx context.Context, id string) error {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	binaryUUID := primitive.Binary{
		Subtype: 0x00,
		Data:    uuid[:],
	}

	_, err = pr.Collection.DeleteOne(ctx, bson.M{"_id": binaryUUID})

	if err != nil {
		return err
	}

	return nil
}

func (pr *ProductRepository) GetProducts(ctx context.Context, categoryId string, page, limit int) ([]domain.Product, error) {
	var products []domain.Product
	opts := options.Find().SetSkip(int64((page - 1) * limit)).SetLimit(int64(limit))
	filter := bson.M{}

	if categoryId != "" {
		uuid, err := uuid.Parse(categoryId)
		if err != nil {
			return nil, err
		}

		binaryUUID := primitive.Binary{
			Subtype: 0x00,
			Data:    uuid[:],
		}

		filter["category_id"] = binaryUUID
	}

	cursor, err := pr.Collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product domain.Product
		if err = cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
