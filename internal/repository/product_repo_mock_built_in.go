package repository

import (
	"context"
)

// MockProductRepo is a mock implementation of the IProductRepo interface
type MockProductRepo struct {
	CreateFn  func(ctx context.Context, product *Product) error
	GetByIDFn func(ctx context.Context, id string) (*Product, error)
	UpdateFn  func(ctx context.Context, product *Product) error
	DeleteFn  func(ctx context.Context, id string) error
	ListFn    func(ctx context.Context) ([]*Product, error)
}

func (m *MockProductRepo) Create(ctx context.Context, product *Product) error {
	if m.CreateFn != nil {
		return m.CreateFn(ctx, product)
	}
	return nil
}

func (m *MockProductRepo) GetByID(ctx context.Context, id string) (*Product, error) {
	if m.GetByIDFn != nil {
		return m.GetByIDFn(ctx, id)
	}
	return nil, nil
}

func (m *MockProductRepo) Update(ctx context.Context, product *Product) error {
	if m.UpdateFn != nil {
		return m.UpdateFn(ctx, product)
	}
	return nil
}

func (m *MockProductRepo) Delete(ctx context.Context, id string) error {
	if m.DeleteFn != nil {
		return m.DeleteFn(ctx, id)
	}
	return nil
}

func (m *MockProductRepo) List(ctx context.Context) ([]*Product, error) {
	if m.ListFn != nil {
		return m.ListFn(ctx)
	}
	return nil, nil
}
