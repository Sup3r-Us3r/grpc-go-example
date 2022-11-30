package server

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/Sup3r-Us3r/go-grpc/grpc/pb"
	"github.com/Sup3r-Us3r/go-grpc/grpc/service"
	"github.com/Sup3r-Us3r/go-grpc/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

var ProductList = model.NewProducts()

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	rootDirectory, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	pemClientCA, err := ioutil.ReadFile(
		filepath.Join(rootDirectory, "grpc", "cert", "ca-cert.pem"),
	)

	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()

	if !certPool.AppendCertsFromPEM(pemClientCA) {
		return nil, fmt.Errorf("failed to add client CA's certificate")
	}

	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair(
		filepath.Join(rootDirectory, "grpc", "cert", "server-cert.pem"),
		filepath.Join(rootDirectory, "grpc", "cert", "server-key.pem"),
	)

	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	return credentials.NewTLS(config), nil
}

func SetupGrpcServer() {
	tlsCredentials, err := loadTLSCredentials()

	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	listener, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(tlsCredentials))
	reflection.Register(grpcServer)

	productService := service.NewProductGrpcService()
	productService.Products = ProductList
	pb.RegisterProductServiceServer(grpcServer, productService)

	fmt.Println("GRPC IS RUNNING ON PORT 50051")

	grpcServer.Serve(listener)
}

// gRPC server without security
// func SetupGrpcServer() {
// 	listener, err := net.Listen("tcp", "localhost:50051")

// 	if err != nil {
// 		panic(err)
// 	}

// 	grpcServer := grpc.NewServer()
// 	reflection.Register(grpcServer)

// 	productService := service.NewProductGrpcService()
// 	productService.Products = ProductList
// 	pb.RegisterProductServiceServer(grpcServer, productService)

// 	fmt.Println("GRPC IS RUNNING ON PORT 50051")

// 	grpcServer.Serve(listener)
// }
