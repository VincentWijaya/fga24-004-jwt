package helpers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const SECRETKEY = "s3CrEtK3Y!"

func GenerateToken(id uint, email string) (string, error) {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	res, err := jwt.SignedString([]byte(SECRETKEY))

	return res, err
}

func VerifyToken(ctx *gin.Context) (interface{}, error) {
	err := errors.New("please login to get the token")
	auth := ctx.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(auth, "Bearer")

	if !bearer {
		return nil, err
	}

	tokenStr := strings.Split(auth, "Bearer ")[1]

	fmt.Println(tokenStr)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRETKEY), nil
	})

	switch {
	case token.Valid:
		fmt.Println("You look nice today")
	case errors.Is(err, jwt.ErrTokenMalformed):
		fmt.Println("That's not even a token")
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		// Invalid signature
		fmt.Println("Invalid signature")
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		// Token is either expired or not active yet
		fmt.Println("Timing is everything")
	default:
		fmt.Println("Couldn't handle this token:", err)
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok {
		fmt.Println("2 ====> ", ok)

		return nil, err
	}

	return token.Claims.(jwt.MapClaims), nil
}
