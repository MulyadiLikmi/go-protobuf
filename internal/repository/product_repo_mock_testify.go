package repository

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// MockProductRepository is a mock implementation of ProductRepository for testing
type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) Create(ctx context.Context, product *Product) error {
	args := m.Called(ctx, product)
	return args.Error(0)
}

func (m *MockProductRepository) GetByID(ctx context.Context, id string) (*Product, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*Product), args.Error(1)
}

func (m *MockProductRepository) Update(ctx context.Context, product *Product) error {
	args := m.Called(ctx, product)
	return args.Error(0)
}

func (m *MockProductRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockProductRepository) List(ctx context.Context) ([]*Product, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*Product), args.Error(1)
}
