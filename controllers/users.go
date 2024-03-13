package controllers

import (
	"belajar-jwt/database"
	"belajar-jwt/helpers"
	"belajar-jwt/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CreateUser(ctx *gin.Context) {
	request := &models.CreateUserRequest{}

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

	hashedPassword := helpers.HashPassword(request.Password)

	userData := &models.User{
		FullName: request.FullName,
		Email:    request.Email,
		Password: hashedPassword,
	}

	db := database.GetPostgres()

	err = db.Create(userData).Error
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(500, map[string]interface{}{
			"message": "Failed to create user",
			"detail":  err,
		})
		return
	}

	ctx.JSON(200, request)
}

func Login(ctx *gin.Context) {
	request := &models.LoginRequest{}

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

	user := &models.User{Email: request.Email}
	err = db.First(user).Error
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(403, map[string]interface{}{
			"message": "Wrong email/password",
		})
		return
	}

	result := helpers.ValidateHashPassword(request.Password, user.Password)
	if !result {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(403, map[string]interface{}{
			"message": "Wrong email/password",
		})
		return
	}

	jwt, err := helpers.GenerateToken(user.ID, user.Email)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(500, map[string]interface{}{
			"message": "Failed to create jwt",
		})
		return
	}

	ctx.JSON(200, map[string]interface{}{
		"token": jwt,
	})
}
