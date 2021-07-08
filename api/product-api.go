package api

import (
	"github.com/babay15/product-app/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductApi struct {
	loginController controller.LoginController
	productController controller.ProductController
}

func NewProductAPI(loginController controller.LoginController, productController controller.ProductController) *ProductApi {
	return &ProductApi{
		loginController: loginController,
		productController: productController,
	}
}

func (api *ProductApi) Authenticate(ctx *gin.Context){
	token := api.loginController.Login(ctx)
	if token != "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":http.StatusOK,
			"token":token,
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":http.StatusUnauthorized,
			"message":"You are not authorized to access this service",
		})
	}
}

func (api *ProductApi) FindAllProducts(ctx *gin.Context){
	ctx.JSON(200, api.productController.FindAllProducts())
}

func (api *ProductApi) SaveProduct(ctx *gin.Context){
	err := api.productController.SaveProduct(ctx)
	if err!=nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code" : http.StatusBadRequest,
			"error": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"code" : http.StatusOK,
			"message" : "Product successfully saved",
		})
	}
}

func (api *ProductApi) UpdateProduct(ctx *gin.Context){
	err := api.productController.UpdateProduct(ctx)
	if err!=nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code" : http.StatusBadRequest,
			"error": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"code" : http.StatusOK,
			"message" : "Product successfully updated",
		})
	}
}

func (api *ProductApi) DeleteProduct(ctx *gin.Context){
	err := api.productController.DeleteProduct(ctx)
	if err!=nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code" : http.StatusBadRequest,
			"error": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"code" : http.StatusOK,
			"message" : "Product successfully deleted",
		})
	}
}

func (api *ProductApi) FindAllBrands(ctx *gin.Context){
	ctx.JSON(200, api.productController.FindAllBrands())
}

func (api *ProductApi) SaveBrand(ctx *gin.Context){
	err := api.productController.SaveBrand(ctx)
	if err!=nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code" : http.StatusBadRequest,
			"error": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"code" : http.StatusOK,
			"message" : "Brand successfully saved",
		})
	}
}

func (api *ProductApi) UpdateBrand(ctx *gin.Context){
	err := api.productController.UpdateBrand(ctx)
	if err!=nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code" : http.StatusBadRequest,
			"error": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"code" : http.StatusOK,
			"message" : "Brand successfully updated",
		})
	}
}

func (api *ProductApi) DeleteBrand(ctx *gin.Context){
	err := api.productController.DeleteBrand(ctx)
	if err!=nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code" : http.StatusBadRequest,
			"error": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"code" : http.StatusOK,
			"message" : "Brand successfully deleted",
		})
	}
}

func (api *ProductApi) ViewProducts(ctx *gin.Context){
	api.productController.ShowProduct(ctx)
}