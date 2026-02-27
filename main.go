package main

import (
	"os"

	"challenge/internal/delivery/http"
	"challenge/internal/repository/memory"
	"challenge/internal/usecase"
)

func main() {
	repo := memory.NewBookMemoryRepository()
	bookUC := usecase.NewBookUsecase(repo)
	authUC := &usecase.AuthUsecase{}

	port := os.Getenv("PORT")
	if port == "" {
		port = "1234"
	}

	r := http.NewRouter(bookUC, authUC)
	r.Run(":" + port)
}
