package database

import (
	"belajar-jwt/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func NewPostgres() {
	fmt.Println("Init db")
	dsn := "host=localhost user=postgres password=test123456 dbname=belajar-jwt port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println("Fisnihed db")
	if err != nil {
		panic(fmt.Sprintf("Failed to init connection to database: %v", err))
	}

	db.Debug().AutoMigrate(&models.User{}, &models.Product{})
}

func GetPostgres() *gorm.DB {
	return db
}
