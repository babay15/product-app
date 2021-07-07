package controller

import (
	"github.com/babay15/product-app/entity"
	"github.com/babay15/product-app/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductController interface {
	FindAll() []entity.Product
	Save(ctx *gin.Context) error
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

func (c *controller) FindAll() []entity.Product {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error{
	var product entity.Product

	if err:= ctx.ShouldBindJSON(&product); err != nil {
		return err
	}

	ctx.BindJSON(&product)
	c.service.Save(product)
	return nil
}

func (c *controller) ShowProduct(ctx *gin.Context) {
	products := c.service.FindAll()
	data := gin.H{
		"title" : "Product List",
		"products" : products,
	}

	ctx.HTML(http.StatusOK, "index.html", data)
}