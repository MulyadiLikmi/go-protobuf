package main

import (
	"context"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"go-protobuf/internal/product_grpc"
	"go-protobuf/internal/repository"
	"go-protobuf/internal/service"
	"go-protobuf/proto"
)

func main() {
	// MongoDB connection setup
	mongoURI := "mongodb://localhost:27017" // Replace with your MongoDB URI
	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Check connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB")

	// Create the repository and service instances
	productRepository := repository.NewMongoProductRepository(client.Database("product_protobuf"))
	productService := service.NewProductService(productRepository)

	// Set up gRPC server
	grpcServer := grpc.NewServer()
	proto.RegisterProductServiceServer(grpcServer, product_grpc.NewProductGRPCServer(productService))

	// Register reflection service on gRPC server (for debugging)
	reflection.Register(grpcServer)

	// Listen and serve
	lis, err := net.Listen("tcp", ":50051") // Replace with your desired port
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	log.Println("Starting gRPC server on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}

// Generate file
// protoc --go_out=. --go-grpc_out=. proto/product.proto

// GET All Products
// grpcurl -d "{}" -plaintext localhost:50051 proto.ProductService/ListProducts

// GET Product By Id
// grpcurl -d "{\"id\": \"1\"}" -plaintext localhost:50051 proto.ProductService/GetProductByID

// Add Product
// grpcurl -d "{\"id\": \"1\", \"name\": \"sample1\", \"stock\": 10}" -plaintext localhost:50051 proto.ProductService/CreateProduct

// Update Product
// grpcurl -d "{\"id\": \"1\", \"name\": \"updated1\", \"stock\": 20}" -plaintext localhost:50051 proto.ProductService/UpdateProduct

// Delete Product
// grpcurl -d "{\"id\": \"1\"}" -plaintext localhost:50051 proto.ProductService/DeleteProduct

// Testing
// go test -v ./test_testify
// go test -v ./test_builtin
