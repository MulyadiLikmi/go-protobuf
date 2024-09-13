package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoProductRepository struct {
	collection *mongo.Collection
}

func NewMongoProductRepository(db *mongo.Database) *MongoProductRepository {
	return &MongoProductRepository{
		collection: db.Collection("products"),
	}
}

func (r *MongoProductRepository) Create(ctx context.Context, product *Product) error {
	_, err := r.collection.InsertOne(ctx, product)
	return err
}

func (r *MongoProductRepository) GetByID(ctx context.Context, id string) (*Product, error) {
	var product Product
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	return &product, err
}

func (r *MongoProductRepository) Update(ctx context.Context, product *Product) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": product.ID},
		bson.M{"$set": product},
	)
	return err
}

func (r *MongoProductRepository) Delete(ctx context.Context, id string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *MongoProductRepository) List(ctx context.Context) ([]*Product, error) {
	var products []*Product
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &products); err != nil {
		return nil, err
	}
	return products, nil
}
