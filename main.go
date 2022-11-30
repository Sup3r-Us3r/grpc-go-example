package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sup3r-Us3r/go-grpc/grpc/server"
)

func handleProducts(w http.ResponseWriter, _ *http.Request) {
	productListJson, err := json.Marshal(server.ProductList)

	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(productListJson))
}

func setupHttpServer() {
	fmt.Println("HTTP SERVER IS RUNNING ON PORT 8080")

	http.HandleFunc("/products", handleProducts)
	http.ListenAndServe(":8080", nil)
}

func main() {
	go server.SetupGrpcServer()
	setupHttpServer()
}
