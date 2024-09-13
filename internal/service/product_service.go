package service

import (
	"context"
	"go-protobuf/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(ctx context.Context, product *repository.Product) error {
	return s.repo.Create(ctx, product)
}

func (s *ProductService) GetProductByID(ctx context.Context, id string) (*repository.Product, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ProductService) UpdateProduct(ctx context.Context, product *repository.Product) error {
	return s.repo.Update(ctx, product)
}

func (s *ProductService) DeleteProduct(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *ProductService) ListProducts(ctx context.Context) ([]*repository.Product, error) {
	return s.repo.List(ctx)
}
