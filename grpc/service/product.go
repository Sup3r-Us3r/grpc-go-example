package service

import (
	"context"

	"github.com/Sup3r-Us3r/go-grpc/grpc/pb"
	"github.com/Sup3r-Us3r/go-grpc/model"
	"github.com/golang/protobuf/ptypes/empty"
)

type ProductGrpcService struct {
	pb.UnimplementedProductServiceServer
	Products *model.Products
}

func NewProductGrpcService() *ProductGrpcService {
	return &ProductGrpcService{}
}

func (p *ProductGrpcService) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product := model.NewProduct()
	product.Name = in.GetName()

	p.Products.Add(product)

	return &pb.CreateProductResponse{
		Product: &pb.Product{
			Id:   product.Id,
			Name: product.Name,
		},
	}, nil
}

func (p *ProductGrpcService) ListProducts(ctx context.Context, in *empty.Empty) (*pb.ListProductsResponse, error) {
	var productItems []*pb.Product

	for _, product := range p.Products.Products {
		productItems = append(
			productItems,
			&pb.Product{Id: product.Id, Name: product.Name},
		)
	}

	return &pb.ListProductsResponse{
		Products: productItems,
	}, nil
}
