package main

import (
	"burakyucel/test/controllers"
	"burakyucel/test/middlewares"
	"burakyucel/test/repositories"
	"burakyucel/test/services"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	productRepository repositories.ProductRepository = repositories.NewVideoRepository()
	productService    services.ProductService        = services.New(productRepository)
	productController controllers.ProductController  = controllers.New(productService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	defer productRepository.CloseDB()

	setupLogOutput()

	server := gin.New()

	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	server.Use(
		gin.Recovery(),
		middlewares.Logger(),
	)

	apiRoutes := server.Group("/api/v1")
	{
		apiRoutes.GET("/products", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, productController.FindAll())
		})

		apiRoutes.POST("/products", func(ctx *gin.Context) {
			err := productController.Store(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Success",
				})
			}

		})

		apiRoutes.PUT("/products/:id", func(ctx *gin.Context) {
			err := productController.Update(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Success",
				})
			}

		})

		apiRoutes.DELETE("/products/:id", func(ctx *gin.Context) {
			err := productController.Delete(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Success",
				})
			}

		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/products", productController.ShowAll)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	server.Run(":" + port)
}
