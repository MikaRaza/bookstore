package routes

import (
	"example/bookstore/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterBookRoutes(r *gin.Engine) {
	r.GET("/books", handlers.GetBooks)
	r.POST("/books", handlers.PostBooks)
	r.GET("/books/:id", handlers.GetBookByID)
}
