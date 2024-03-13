package router

import (
	"belajar-jwt/controllers"
	"belajar-jwt/middlewares"

	"github.com/gin-gonic/gin"
)

func NewRouter() {
	r := gin.Default()

	r.GET("", rootHandler)

	users := r.Group("/users")
	users.POST("/", controllers.CreateUser)
	users.POST("/login", controllers.Login)

	products := r.Group("/products")
	{
		products.Use(middlewares.Authentication())
		products.POST("/", controllers.CreateProduct)
	}

	r.Run(":8080")
}

func rootHandler(ctx *gin.Context) {
	ctx.JSON(200, map[string]interface{}{
		"message": "Service is up and running!",
	})
}
