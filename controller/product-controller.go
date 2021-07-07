package controller

import (
	"github.com/babay15/product-app/entity"
	"github.com/babay15/product-app/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductController interface {
	SaveProduct(ctx *gin.Context) error
	UpdateProduct(ctx *gin.Context) error
	DeleteProduct(ctx *gin.Context) error
	FindAllProducts() []entity.Product
	SaveBrand(ctx *gin.Context) error
	UpdateBrand(ctx *gin.Context) error
	DeleteBrand(ctx *gin.Context) error
	FindAllBrands() []entity.Brand

	ShowProduct(ctx *gin.Context)
}

type controller struct {
	service service.ProductService
}

func New(productService service.ProductService) ProductController{
	return &controller{
		service: productService,
	}
}

func (c *controller) SaveProduct(ctx *gin.Context) error{
	var product entity.Product

	if err:= ctx.ShouldBindJSON(&product); err != nil {
		return err
	}

	ctx.BindJSON(&product)
	c.service.SaveProduct(product)
	return nil
}

func (c *controller) UpdateProduct(ctx *gin.Context) error{
	var product entity.Product
	id := ctx.Param("id")
	if err:= ctx.ShouldBindJSON(&product); err != nil {
		return err
	}

	product.ID = id
	ctx.BindJSON(&product)
	c.service.UpdateProduct(product)
	return nil
}

func (c *controller) DeleteProduct(ctx *gin.Context) error{
	var product entity.Product
	id := ctx.Param("id")
	if err:= ctx.ShouldBindJSON(&product); err != nil {
		return err
	}

	product.ID = id
	ctx.BindJSON(&product)
	c.service.DeleteProduct(product)
	return nil
}

func (c *controller) FindAllProducts() []entity.Product {
	return c.service.FindAllProduct()
}

func (c *controller) SaveBrand(ctx *gin.Context) error{
	var brand entity.Brand

	if err:= ctx.ShouldBindJSON(&brand); err != nil {
		return err
	}

	ctx.BindJSON(&brand)
	c.service.SaveBrand(brand)
	return nil
}

func (c *controller) UpdateBrand(ctx *gin.Context) error{
	var brand entity.Brand
	id := ctx.Param("id")
	if err:= ctx.ShouldBindJSON(&brand); err != nil {
		return err
	}

	brand.BrandID = id
	ctx.BindJSON(&brand)
	c.service.UpdateBrand(brand)
	return nil
}

func (c *controller) DeleteBrand(ctx *gin.Context) error{
	var brand entity.Brand
	id := ctx.Param("id")
	if err:= ctx.ShouldBindJSON(&brand); err != nil {
		return err
	}

	brand.BrandID = id
	ctx.BindJSON(&brand)
	c.service.DeleteBrand(brand)
	return nil
}

func (c *controller) FindAllBrands() []entity.Brand {
	return c.service.FindAllBrand()
}

func (c *controller) ShowProduct(ctx *gin.Context) {
	products := c.service.FindAllProduct()
	data := gin.H{
		"title" : "Product List",
		"products" : products,
	}

	ctx.HTML(http.StatusOK, "index.html", data)
}