package main

import (
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3307)/pustakaapi?charset=utf8mb4&parseTime=True&locLocal"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection error")
	}

	db.AutoMigrate(book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	routerV1 := router.Group("/v1")

	routerV1.GET("/", handler.RootHandler)

	routerV1.GET("/hello", handler.HelloHandler)

	routerV1.GET("/book/:id/:title", handler.BookHandler)

	routerV1.GET("/query", handler.QueryHandler)

	routerV1.GET("/books", bookHandler.GetBooks)
	routerV1.GET("/books/:id", bookHandler.GetBook)
	routerV1.POST("/books", bookHandler.PostBookHandler)
	routerV1.PUT("/books/:id", bookHandler.UpdateBookHandler)
	routerV1.DELETE("/books/:id", bookHandler.DeleteBook)

	router.Run(":3030")
}
