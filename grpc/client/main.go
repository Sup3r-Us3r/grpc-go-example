package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/Sup3r-Us3r/go-grpc/grpc/pb"
	"github.com/golang/protobuf/ptypes/empty"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	connection *grpc.ClientConn
)

func init() {
	tlsCredentials, err := loadTLSCredentials()

	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	// connection, err := grpc.Dial("0.0.0.0:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	connection, err = grpc.Dial("0.0.0.0:50051", grpc.WithTransportCredentials(tlsCredentials))

	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	rootDirectory, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := os.ReadFile(
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
	client := pb.NewProductServiceClient(connection)

	request := &empty.Empty{}
	response, err := client.ListProducts(context.Background(), request)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(response)
}

func beatsPerMinute() {
	client := pb.NewSmartwatchServiceClient(connection)

	request := &pb.BeatsPerMinuteRequest{
		Uuid: uuid.NewV4().String(),
	}

	stream, err := client.BeatsPerMinute(context.Background(), request)

	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	done := make(chan bool)
	defer close(done)

	go func() {
		for {
			response, err := stream.Recv()

			if err == io.EOF {
				done <- true
				return
			}

			if err != nil {
				log.Fatalf("cannot receive %v", err)
			}

			log.Printf("stream data received: %s", response.String())
		}
	}()

	<-done

	fmt.Println("Finished stream")
}

func main() {
	listProducts()
	fmt.Println()
	createProduct()
	fmt.Println()
	beatsPerMinute()
}
