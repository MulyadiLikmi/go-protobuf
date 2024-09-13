package product_grpc

import (
	"context"
	"go-protobuf/internal/repository"
	"go-protobuf/internal/service"
	"go-protobuf/proto"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ProductGRPCServer struct {
	proto.UnimplementedProductServiceServer
	service *service.ProductService
}

func NewProductGRPCServer(svc *service.ProductService) *ProductGRPCServer {
	return &ProductGRPCServer{service: svc}
}

func (s *ProductGRPCServer) CreateProduct(ctx context.Context, req *proto.Product) (*proto.ProductResponse, error) {
	product := &repository.Product{
		ID:    req.Id, // Use the ID from the request directly
		Name:  req.Name,
		Stock: req.Stock,
	}

	err := s.service.CreateProduct(ctx, product)
	return &proto.ProductResponse{
		Success: err == nil,
		Message: "Created",
	}, err
}

func (s *ProductGRPCServer) GetProductByID(ctx context.Context, req *proto.Product) (*proto.Product, error) {
	product, err := s.service.GetProductByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &proto.Product{
		Id:    product.ID,
		Name:  product.Name,
		Stock: product.Stock,
	}, nil
}

func (s *ProductGRPCServer) UpdateProduct(ctx context.Context, req *proto.Product) (*proto.ProductResponse, error) {
	product := &repository.Product{
		ID:    req.Id,
		Name:  req.Name,
		Stock: req.Stock,
	}
	err := s.service.UpdateProduct(ctx, product)
	return &proto.ProductResponse{
		Success: err == nil,
		Message: "Updated",
	}, err
}

func (s *ProductGRPCServer) DeleteProduct(ctx context.Context, req *proto.Product) (*proto.ProductResponse, error) {
	err := s.service.DeleteProduct(ctx, req.Id)
	return &proto.ProductResponse{
		Success: err == nil,
		Message: "Deleted",
	}, err
}

func (s *ProductGRPCServer) ListProducts(ctx context.Context, _ *emptypb.Empty) (*proto.ProductList, error) {
	products, err := s.service.ListProducts(ctx)
	if err != nil {
		return nil, err
	}
	var protoProducts []*proto.Product
	for _, product := range products {
		protoProducts = append(protoProducts, &proto.Product{
			Id:    product.ID,
			Name:  product.Name,
			Stock: product.Stock,
		})
	}
	return &proto.ProductList{Products: protoProducts}, nil
}
