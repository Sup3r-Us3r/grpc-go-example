package model

import uuid "github.com/satori/go.uuid"

type Product struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Products struct {
	Products []Product `json:"products"`
}

func NewProduct() *Product {
	return &Product{
		Id: uuid.NewV4().String(),
	}
}

func NewProducts() *Products {
	return &Products{}
}

func (p *Products) Add(product *Product) {
	p.Products = append(p.Products, *product)
}
