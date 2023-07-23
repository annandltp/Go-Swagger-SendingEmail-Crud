package main

import (
	"gin/config"
	"gin/controller"
	"net/http"

	_ "gin/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Trial Class Mini Ecommerce
// @version         1.0
// @description     Dokomentasi REST API project Mini Ecommerce II

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @host      localhost:8000
func main() {
	s := gin.Default()
	config.DBConnect()

	s.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "hello world")
	})

	// list products
	s.GET("/products", controller.HandlerGetProduct)

	// create order
	s.POST("/order", controller.HandlerPostOrder)

	// list order
	s.GET("/orders", controller.HandlerGetOrder)

	s.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s.Run(":8000")
}
