package main

import (
	"latihan3-gin-gorm/handlers"
	"latihan3-gin-gorm/models"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user:postgres password=postgres dbname=userr16 port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	log.Println("DSN:", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Koneksi Error", err)
	}

	db.AutoMigrate(&models.User{})

	r := gin.Default()

	r.GET("/users", handlers.GetUser(db))
	r.POST("/users", handlers.CreateUser(db))

	r.Run(":8080")
}
