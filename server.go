package main

import (
	"github.com/babay15/product-app/controller"
	"github.com/babay15/product-app/middleware"
	"github.com/babay15/product-app/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
	"io"
	"net/http"
	"os"
)

var (
	productService service.ProductService = service.New()
	productController controller.ProductController = controller.New(productService)
)

func createLogOutput() {
	file, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}

func main() {

	createLogOutput()

	//Create Server
	server := gin.Default()

	//Route the styling
	server.Static("/css", "./template/css")

	//Route the html
	server.LoadHTMLGlob("template/*.html")

	//Server Handler
	server.Use(gin.Recovery(), middleware.Logger(), gindump.Dump())

	//To group the server into a path
	apiRoutes := server.Group("/api", middleware.BasicAuth())
	{
		//Calling the controller
		apiRoutes.GET("/products", func(ctx *gin.Context) {
			ctx.JSON(200, productController.FindAll())
		})

		//Calling the controller
		apiRoutes.POST("/save", func(ctx *gin.Context) {
			err := productController.Save(ctx)
			if err!=nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"code" : http.StatusBadRequest,
					"error": err.Error(),
				})
			} else {
				ctx.JSON(200, gin.H{
					"code" : http.StatusOK,
					"message" : "Product succesfully saved",
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
