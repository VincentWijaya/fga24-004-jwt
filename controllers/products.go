package controllers

import (
	"belajar-jwt/database"
	"belajar-jwt/helpers"
	"belajar-jwt/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt/v5"
)

func CreateProduct(ctx *gin.Context) {
	request := &models.CreateProdcuctRequest{}

	err := ctx.ShouldBindBodyWith(request, binding.JSON)
	if err != nil {
		ctx.AbortWithStatusJSON(400, map[string]interface{}{
			"message": "Invalid request",
		})
		return
	}

	validator := helpers.NewValidator()

	err = validator.Validate(request)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(400, map[string]interface{}{
			"message": "Invalid request format",
			"detail":  err.Error(),
		})
		return
	}

	db := database.GetPostgres()

	user := ctx.MustGet("userData").(jwt.MapClaims)
	fmt.Printf("User data: %+v", user)
	product := &models.Product{
		Title:       request.Title,
		Description: request.Description,
		UserID:      uint(user["id"].(float64)),
	}
	err = db.Create(product).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create product",
		})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}
