package repository

import (
	"context"
)

type Product struct {
	ID    string `bson:"_id,omitempty"`
	Name  string `bson:"name"`
	Stock int32  `bson:"stock"`
}

type IProductRepo interface {
	Create(ctx context.Context, product *Product) error
	GetByID(ctx context.Context, id string) (*Product, error)
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*Product, error)
}
