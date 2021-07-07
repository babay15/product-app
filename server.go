package main

import (
	"github.com/babay15/product-app/controller"
	"github.com/babay15/product-app/middleware"
	repository2 "github.com/babay15/product-app/repository"
	"github.com/babay15/product-app/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
	"io"
	"net/http"
	"os"
)

var (
	repository repository2.Repository = repository2.NewRepository()
	productService service.ProductService = service.New(repository)
	loginService service.LoginService = service.NewLoginService()
	jwtService service.JWTService = service.NewJWTService()

	productController controller.ProductController = controller.New(productService)
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

func createLogOutput() {
	file, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}

func main() {
	defer repository.CloseDB()

	createLogOutput()

	//Create Server
	server := gin.Default()

	//Route the styling
	server.Static("/css", "./template/css")

	//Route the html
	server.LoadHTMLGlob("template/*.html")

	//Server Handler
	server.Use(gin.Recovery(), middleware.Logger(), gindump.Dump())

	server.POST("/login", func(context *gin.Context) {
		token := loginController.Login(context)
		if token != "" {
			context.JSON(http.StatusOK, gin.H{
				"code":http.StatusOK,
				"token":token,
			})
		} else {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code":http.StatusUnauthorized,
				"message":"You are not authorized to access this service",
			})
		}
	})

	//To group the server into a path
	productRoutes := server.Group("/api/products", middleware.JwtAuth())
	{
		//Calling the controller
		productRoutes.GET("/all", func(ctx *gin.Context) {
			ctx.JSON(200, productController.FindAllProducts())
		})

		//Calling the controller
		productRoutes.POST("/save", func(ctx *gin.Context) {
			err := productController.SaveProduct(ctx)
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
		})

		//Calling the controller
		productRoutes.POST("/update:id", func(ctx *gin.Context) {
			err := productController.UpdateProduct(ctx)
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
		})

		//Calling the controller
		productRoutes.POST("/delete:id", func(ctx *gin.Context) {
			err := productController.DeleteProduct(ctx)
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
		})
	}

	brandRoutes := server.Group("/api/brands", middleware.JwtAuth())
	{
		//Calling the controller
		brandRoutes.GET("/all", func(ctx *gin.Context) {
			ctx.JSON(200, productController.FindAllBrands())
		})

		//Calling the controller
		brandRoutes.POST("/save", func(ctx *gin.Context) {
			err := productController.SaveBrand(ctx)
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
		})

		//Calling the controller
		brandRoutes.POST("/update:id", func(ctx *gin.Context) {
			err := productController.UpdateBrand(ctx)
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
		})

		//Calling the controller
		brandRoutes.POST("/delete:id", func(ctx *gin.Context) {
			err := productController.DeleteBrand(ctx)
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
		})
	}

	viewRoutes := server.Group("/view")
	{
		//Calling the controller
		viewRoutes.GET(
			"/products",
			productController.ShowProduct,
		)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	server.Run(":" + port)
}
