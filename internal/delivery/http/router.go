package http

import (
	"challenge/internal/delivery/http/handler"
	"challenge/internal/usecase"

	"github.com/gin-gonic/gin"
)

func NewRouter(bookUC *usecase.BookUsecase, authUC *usecase.AuthUsecase) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", handler.Ping)
	r.POST("/echo", handler.Echo)

	bookHandler := handler.NewBookHandler(bookUC)

	r.POST("/books", bookHandler.Create)
	r.GET("/books", bookHandler.GetAll)
	r.GET("/books/:id", bookHandler.GetByID)
	r.PUT("/books/:id", bookHandler.Update)
	r.DELETE("/books/:id", bookHandler.Delete)

	r.POST("/auth/token", func(c *gin.Context) {
		token, _ := authUC.GenerateToken("admin")
		c.JSON(200, gin.H{"token": token})
	})

	return r
}
