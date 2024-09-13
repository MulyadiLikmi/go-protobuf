package test_service

import (
	"context"
	"go-protobuf/internal/repository"
	"go-protobuf/internal/service"
	"testing"
)

func TestProductService_CreateProduct(t *testing.T) {
	// Assuming you have a mock repository with a Create method
	mockRepo := &repository.MockProductRepository{}
	mockRepo.On("Create", context.Background(), &repository.Product{ID: "1", Name: "Product 1", Stock: 10}).Return(nil)

	productService := service.NewProductService(mockRepo)

	err := productService.CreateProduct(context.Background(), &repository.Product{ID: "1", Name: "Product 1", Stock: 10})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestProductService_GetProductByID(t *testing.T) {
	mockRepo := &repository.MockProductRepository{}
	mockRepo.On("GetByID", context.Background(), "1").Return(&repository.Product{ID: "1", Name: "Product 1", Stock: 10}, nil)

	productService := service.NewProductService(mockRepo)
	result, err := productService.GetProductByID(context.Background(), "1")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.ID != "1" || result.Name != "Product 1" || result.Stock != 10 {
		t.Fatalf("expected product %v, got %v", &repository.Product{ID: "1", Name: "Product 1", Stock: 10}, result)
	}
}

func TestProductService_UpdateProduct(t *testing.T) {
	mockRepo := &repository.MockProductRepository{}
	mockRepo.On("Update", context.Background(), &repository.Product{ID: "1", Name: "Product 1", Stock: 10}).Return(nil)

	productService := service.NewProductService(mockRepo)

	err := productService.UpdateProduct(context.Background(), &repository.Product{ID: "1", Name: "Product 1", Stock: 10})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestProductService_DeleteProduct(t *testing.T) {
	mockRepo := &repository.MockProductRepository{}
	mockRepo.On("Delete", context.Background(), "1").Return(nil)

	productService := service.NewProductService(mockRepo)

	err := productService.DeleteProduct(context.Background(), "1")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestProductService_ListProducts(t *testing.T) {
	mockRepo := &repository.MockProductRepository{}
	mockRepo.On("List", context.Background()).Return([]*repository.Product{
		{ID: "1", Name: "Product 1", Stock: 10},
		{ID: "2", Name: "Product 2", Stock: 20},
	}, nil)

	productService := service.NewProductService(mockRepo)
	result, err := productService.ListProducts(context.Background())

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	expected := []*repository.Product{
		{ID: "1", Name: "Product 1", Stock: 10},
		{ID: "2", Name: "Product 2", Stock: 20},
	}
	if len(result) != len(expected) {
		t.Fatalf("expected %v products, got %v", len(expected), len(result))
	}
	for i, product := range result {
		if product.ID != expected[i].ID || product.Name != expected[i].Name || product.Stock != expected[i].Stock {
			t.Fatalf("expected product %v, got %v", expected[i], product)
		}
	}
}
