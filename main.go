package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/Sup3r-Us3r/go-grpc/grpc/pb"
	"github.com/Sup3r-Us3r/go-grpc/grpc/service"
	"github.com/Sup3r-Us3r/go-grpc/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var ProductList = model.NewProducts()

func handleProducts(w http.ResponseWriter, _ *http.Request) {
	productListJson, err := json.Marshal(ProductList)

	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-type", "application/json")
	w.Write([]byte(productListJson))
}

func setupGrpcServer() {
	listener, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	productGrpcService := service.NewProductGrpcService()
	productGrpcService.Products = ProductList
	pb.RegisterProductServiceServer(grpcServer, productGrpcService)

	fmt.Println("GRPC IS RUNNING ON PORT 50051")

	grpcServer.Serve(listener)
}

func setupHttpServer() {
	fmt.Println("HTTP SERVER IS RUNNING ON PORT 8080")

	http.HandleFunc("/products", handleProducts)
	http.ListenAndServe(":8080", nil)
}

func main() {
	go setupGrpcServer()
	setupHttpServer()
}
