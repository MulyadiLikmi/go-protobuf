package test_service

import (
	"context"
	"go-protobuf/internal/repository"
	"go-protobuf/internal/service" // Import the service package
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProductService_CreateProduct(t *testing.T) {
	// Initialize the mock repository
	mockRepo := new(repository.MockProductRepository)
	product := &repository.Product{ID: "1", Name: "Product 1", Stock: 10}

	// Set expectations for the mock
	mockRepo.On("Create", mock.Anything, product).Return(nil)

	// Initialize the service
	productService := service.NewProductService(mockRepo)

	// Call the CreateProduct method
	err := productService.CreateProduct(context.Background(), product)

	// Assert the expectations
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductService_GetProductByID(t *testing.T) {
	// Initialize the mock repository
	mockRepo := new(repository.MockProductRepository)
	product := &repository.Product{ID: "1", Name: "Product 1", Stock: 10}

	// Set expectations for the mock
	mockRepo.On("GetByID", mock.Anything, "1").Return(product, nil)

	// Initialize the service
	productService := service.NewProductService(mockRepo)

	// Call the GetProductByID method
	result, err := productService.GetProductByID(context.Background(), "1")

	// Assert the expectations
	assert.NoError(t, err)
	assert.Equal(t, product, result)
	mockRepo.AssertExpectations(t)
}

func TestProductService_UpdateProduct(t *testing.T) {
	// Initialize the mock repository
	mockRepo := new(repository.MockProductRepository)
	product := &repository.Product{ID: "1", Name: "Product 1", Stock: 10}

	// Set expectations for the mock
	mockRepo.On("Update", mock.Anything, product).Return(nil)

	// Initialize the service
	productService := service.NewProductService(mockRepo)

	// Call the UpdateProduct method
	err := productService.UpdateProduct(context.Background(), product)

	// Assert the expectations
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductService_DeleteProduct(t *testing.T) {
	// Initialize the mock repository
	mockRepo := new(repository.MockProductRepository)

	// Set expectations for the mock
	mockRepo.On("Delete", mock.Anything, "1").Return(nil)

	// Initialize the service
	productService := service.NewProductService(mockRepo)

	// Call the DeleteProduct method
	err := productService.DeleteProduct(context.Background(), "1")

	// Assert the expectations
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductService_ListProducts(t *testing.T) {
	// Initialize the mock repository
	mockRepo := new(repository.MockProductRepository)
	products := []*repository.Product{
		{ID: "1", Name: "Product 1", Stock: 10},
		{ID: "2", Name: "Product 2", Stock: 20},
	}

	// Set expectations for the mock
	mockRepo.On("List", mock.Anything).Return(products, nil)

	// Initialize the service
	productService := service.NewProductService(mockRepo)

	// Call the ListProducts method
	result, err := productService.ListProducts(context.Background())

	// Assert the expectations
	assert.NoError(t, err)
	assert.Equal(t, products, result)
	mockRepo.AssertExpectations(t)
}
