package product

import (
	"context"
	pb "ecommerce-grpc/proto/productpb"
)

// Service defines business logic abstraction
type Service interface {
	GetProduct(ctx context.Context, id string) (*Product, error)
	ListProducts(ctx context.Context) ([]*Product, error)
}

type productService struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &productService{repo: repo}
}

func (s *productService) GetProduct(ctx context.Context, id string) (*Product, error) {
	return s.repo.GetByID(id)
}

func (s *productService) ListProducts(ctx context.Context) ([]*Product, error) {
	return s.repo.ListAll()
}

// gRPC adapter
type GRPCHandler struct {
	pb.UnimplementedProductServiceServer
	service Service
}

func NewGRPCHandler(s Service) *GRPCHandler {
	return &GRPCHandler{service: s}
}

func (h *GRPCHandler) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	p, _ := h.service.GetProduct(ctx, req.GetId())
	if p == nil {
		return &pb.GetProductResponse{}, nil
	}
	return &pb.GetProductResponse{
		Product: &pb.Product{
			Id:    p.ID,
			Name:  p.Name,
			Price: p.Price,
		},
	}, nil
}

func (h *GRPCHandler) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	products, _ := h.service.ListProducts(ctx)
	resp := &pb.ListProductsResponse{}
	for _, p := range products {
		resp.Products = append(resp.Products, &pb.Product{
			Id:    p.ID,
			Name:  p.Name,
			Price: p.Price,
		})
	}
	return resp, nil
}
