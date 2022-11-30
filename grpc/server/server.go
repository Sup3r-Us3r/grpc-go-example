package server

import (
	"fmt"
	"net"

	"github.com/Sup3r-Us3r/go-grpc/grpc/pb"
	"github.com/Sup3r-Us3r/go-grpc/grpc/service"
	"github.com/Sup3r-Us3r/go-grpc/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var ProductList = model.NewProducts()

func SetupGrpcServer() {
	listener, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	productService := service.NewProductGrpcService()
	productService.Products = ProductList
	pb.RegisterProductServiceServer(grpcServer, productService)

	fmt.Println("GRPC IS RUNNING ON PORT 50051")

	grpcServer.Serve(listener)
}
