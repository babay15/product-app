package main

import (
	"github.com/babay15/product-app/api"
	"github.com/babay15/product-app/controller"
	"github.com/babay15/product-app/docs"
	"github.com/babay15/product-app/middleware"
	repository2 "github.com/babay15/product-app/repository"
	"github.com/babay15/product-app/service"
	"github.com/gin-gonic/gin"
	"io"
	"os"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

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
	docs.SwaggerInfo.Title = "Product API"
	docs.SwaggerInfo.Description = "Babay15 - Product CRUD API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:5000"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	defer repository.CloseDB()

	createLogOutput()

	//Create Server
	server := gin.Default()

	productAPI := api.NewProductAPI(loginController, productController)

	apiRoutes := server.Group(docs.SwaggerInfo.BasePath)
	{
		login := apiRoutes.Group("/auth")
		{
			login.POST("token", productAPI.Authenticate)
		}

		product := apiRoutes.Group("/products", middleware.JwtAuth())
		{
			product.GET("/all", productAPI.FindAllProducts)
			product.POST("/save", productAPI.SaveProduct)
			product.POST("/update:id", productAPI.UpdateProduct)
			product.POST("/delete:id", productAPI.DeleteProduct)
		}
		brand := apiRoutes.Group("/brand", middleware.JwtAuth())
		{
			brand.GET("/all", productAPI.FindAllBrands)
			brand.POST("/save", productAPI.SaveBrand)
			brand.POST("/update:id", productAPI.UpdateBrand)
			brand.POST("/delete:id", productAPI.DeleteBrand)
		}
	}

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//Route the styling
	server.Static("/css", "./template/css")

	//Route the html
	server.LoadHTMLGlob("template/*.html")

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
