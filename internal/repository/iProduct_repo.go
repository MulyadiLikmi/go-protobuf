package repository

import (
	"context"
)

// Product represents the product entity
type Product struct {
	ID    string `bson:"_id,omitempty"`
	Name  string `bson:"name"`
	Stock int32  `bson:"stock"`
}

// ProductRepository is an interface that defines the methods for product operations
type IProductRepo interface {
	Create(ctx context.Context, product *Product) error
	GetByID(ctx context.Context, id string) (*Product, error)
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*Product, error)
}
