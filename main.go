package main

import (
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"
	"pustaka-api/initializers"
	"pustaka-api/middleware"
	"pustaka-api/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error

	initializers.LoadEnvVariables()
	db, err = initializers.ConnectToDatabase()
	err = initializers.SyncDatabase(db)

	if err != nil {
		log.Fatal("Connection to database failed")
	}
}

func main() {
	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	routerV1 := router.Group("/v1")

	routerV1.POST("/signup", userHandler.Signup)
	routerV1.POST("/login", userHandler.Login)

	routerV1.GET("/", handler.RootHandler)
	routerV1.GET("/hello", handler.HelloHandler)
	routerV1.GET("/book/:id/:title", handler.BookHandler)
	routerV1.GET("/query", handler.QueryHandler)

	routerV1Books := router.Group("/books", middleware.RequireAuth)
	routerV1Books.GET("/books", bookHandler.GetBooks)
	routerV1Books.GET("/books/:id", bookHandler.GetBook)
	routerV1Books.POST("/books", bookHandler.PostBookHandler)
	routerV1Books.PUT("/books/:id", bookHandler.UpdateBookHandler)
	routerV1Books.DELETE("/books/:id", bookHandler.DeleteBook)

	router.Run(":3030")
}
