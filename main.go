package main

import (
	"fmt"
	"lab-gin-crud/book"
	"lab-gin-crud/handlers"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// Gorm
	dsn := "root:admin@tcp(127.0.0.1:3306)/buku_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection ERROR")
	}
	fmt.Println("Database connection succeed")

	// Layering
	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handlers.NewBookHandler(bookService)

	// Migration
	db.AutoMigrate(&book.Book{})

	//Gin Gonic
	router := gin.Default()

	// Versioning
	v1 := router.Group("/v1")

	// Routes
	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBook)
	v1.POST("/books", bookHandler.CreateBook)
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("/books/:id", bookHandler.DeleteBook)

	// default port :8080
	router.Run()
}
