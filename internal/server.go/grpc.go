package server

import (
	pb "ecommerce-grpc/proto/productpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

func RunGRPCServer(addr string, handler pb.ProductServiceServer) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, handler)

	log.Printf("gRPC server running on %s", addr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
