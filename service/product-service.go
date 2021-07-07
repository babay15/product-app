package service

import (
	"github.com/babay15/product-app/entity"
	"github.com/babay15/product-app/repository"
)

type ProductService interface {
	SaveProduct(entity.Product) entity.Product
	UpdateProduct(product entity.Product) entity.Product
	DeleteProduct(product entity.Product) entity.Product
	FindAllProduct() []entity.Product
	SaveBrand(brand entity.Brand) entity.Brand
	UpdateBrand(product entity.Brand) entity.Brand
	DeleteBrand(product entity.Brand) entity.Brand
	FindAllBrand() []entity.Brand
}

type productService struct {
	repository repository.Repository
}

func (p *productService) SaveProduct(product entity.Product) entity.Product {
	p.repository.SaveProduct(product)
	return product
}

func (p *productService) UpdateProduct(product entity.Product) entity.Product {
	p.repository.UpdateProduct(product)
	return product
}

func (p *productService) DeleteProduct(product entity.Product) entity.Product {
	p.repository.DeleteProduct(product)
	return product
}

func (p *productService) FindAllProduct() []entity.Product {
	return p.repository.FindAllProduct()
}

func (p *productService) SaveBrand(brand entity.Brand) entity.Brand {
	p.repository.SaveBrand(brand)
	return brand
}

func (p *productService) UpdateBrand(brand entity.Brand) entity.Brand {
	p.repository.UpdateBrand(brand)
	return brand
}

func (p *productService) DeleteBrand(brand entity.Brand) entity.Brand {
	p.repository.DeleteBrand(brand)
	return brand
}

func (p *productService) FindAllBrand() []entity.Brand {
	return p.repository.FindAllBrand()
}

func New(repository repository.Repository) ProductService {
	return &productService{
		repository: repository,
	}
}

