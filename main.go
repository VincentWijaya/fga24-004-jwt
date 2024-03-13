package main

import (
	"belajar-jwt/database"
	"belajar-jwt/router"
)

func main() {
	database.NewPostgres()
	router.NewRouter()
}
