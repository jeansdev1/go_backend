package main

import (
	"go-api/controller"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	ProductUsecase := usecase.NewProductUseCase()
	// camada de controllers
	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/ping", func (ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"messsage": "pong",
	})		
	})
	server.GET("/products", ProductController.GetProducts)
	server.Run(":8000")
}