package middlewares

import (
	"belajar-jwt/helpers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		userData, err := helpers.VerifyToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token",
			})
			return
		}

		fmt.Println("UYser, ", userData)

		c.Set("userData", userData)
		c.Next()
	}
}
