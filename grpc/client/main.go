package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/Sup3r-Us3r/go-grpc/grpc/pb"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	rootDirectory, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile(
		filepath.Join(rootDirectory, "grpc", "cert", "ca-cert.pem"),
	)

	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()

	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Load client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair(
		filepath.Join(rootDirectory, "grpc", "cert", "client-cert.pem"),
		filepath.Join(rootDirectory, "grpc", "cert", "client-key.pem"),
	)

	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil
}

func createProduct() {
	tlsCredentials, err := loadTLSCredentials()

	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	connection, err := grpc.Dial("0.0.0.0:50051", grpc.WithTransportCredentials(tlsCredentials))

	if err != nil {
		log.Fatal("cannot dial server: ", err)
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

// gRPC client without security
// func createProduct() {
// 	connection, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	client := pb.NewProductServiceClient(connection)

// 	request := &pb.CreateProductRequest{
// 		Name: "Product example",
// 	}
// 	response, err := client.CreateProduct(context.Background(), request)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Print(response)
// }

func listProducts() {
	tlsCredentials, err := loadTLSCredentials()

	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	connection, err := grpc.Dial("0.0.0.0:50051", grpc.WithTransportCredentials(tlsCredentials))

	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	client := pb.NewProductServiceClient(connection)

	request := &empty.Empty{}
	response, err := client.ListProducts(context.Background(), request)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(response)
}

// gRPC client without security
// func listProducts() {
// 	connection, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	client := pb.NewProductServiceClient(connection)

// 	request := &empty.Empty{}
// 	response, err := client.ListProducts(context.Background(), request)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Print(response)
// }

func main() {
	listProducts()
	fmt.Println()
	createProduct()
}
