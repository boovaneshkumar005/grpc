package main

import (
	"ecommerce-grpc/internal/product"
	"ecommerce-grpc/internal/server.go"
)

func main() {
	// Repository
	repo := product.NewInMemoryRepository()

	// Business service
	service := product.NewService(repo)

	// gRPC handler
	handler := product.NewGRPCHandler(service)

	// Start gRPC server
	server.RunGRPCServer(":50051", handler)
}
