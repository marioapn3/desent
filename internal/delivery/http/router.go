package http

import (
	"challenge/internal/delivery/http/handler"
	"challenge/internal/delivery/http/middleware"
	"challenge/internal/usecase"

	"github.com/gin-gonic/gin"
)

func NewRouter(bookUC *usecase.BookUsecase, authUC *usecase.AuthUsecase) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", handler.Ping)
	r.POST("/echo", handler.Echo)

	authHandler := handler.NewAuthHandler(authUC)
	r.POST("/auth/token", authHandler.Login)

	bookHandler := handler.NewBookHandler(bookUC)

	r.POST("/books", bookHandler.Create)
	r.DELETE("/books", bookHandler.DeleteAll)
	r.GET("/books/:id", bookHandler.GetByID)
	r.PUT("/books/:id", bookHandler.Update)
	r.DELETE("/books/:id", bookHandler.Delete)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/books", bookHandler.GetAll)
	}

	return r
}
