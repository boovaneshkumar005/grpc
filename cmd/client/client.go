package main

import (
	"context"
	pb "ecommerce-grpc/proto/productpb"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	// Dial gRPC server
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewProductServiceClient(conn)

	// Test GetProduct
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := client.GetProduct(ctx, &pb.GetProductRequest{Id: "1"})
	if err != nil {
		log.Fatalf("error calling GetProduct: %v", err)
	}

	fmt.Println("GetProduct Response:")
	fmt.Printf("ID: %s, Name: %s, Price: %.2f\n", resp.Product.Id, resp.Product.Name, resp.Product.Price)

	// Test ListProducts
	listResp, err := client.ListProducts(ctx, &pb.ListProductsRequest{})
	if err != nil {
		log.Fatalf("error calling ListProducts: %v", err)
	}

	fmt.Println("\n ListProducts Response:")
	for _, p := range listResp.Products {
		fmt.Printf("ID: %s, Name: %s, Price: %.2f\n", p.Id, p.Name, p.Price)
	}
}
