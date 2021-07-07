package service

import "github.com/babay15/product-app/entity"

type ProductService interface {
	Save(entity.Product) entity.Product
	FindAll() []entity.Product
}

type productService struct {
	product []entity.Product
}

func (p *productService) Save(product entity.Product) entity.Product {
	p.product = append(p.product, product)
	return product
}

func (p *productService) FindAll() []entity.Product {
	return p.product
}

func New() ProductService {
	return &productService{}
}

