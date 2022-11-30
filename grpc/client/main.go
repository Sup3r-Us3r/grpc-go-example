package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Sup3r-Us3r/go-grpc/grpc/pb"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func createProduct() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewProductServiceClient(connection)

	request := &pb.CreateProductRequest{
		Name: "Product example",
	}
	response, err := client.CreateProduct(context.Background(), request)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(response)
}

func listProducts() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewProductServiceClient(connection)

	request := &empty.Empty{}
	response, err := client.ListProducts(context.Background(), request)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(response)
}

func main() {
	listProducts()
	fmt.Println()
	createProduct()
}
