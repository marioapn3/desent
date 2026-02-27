package main

import (
	"challenge/internal/delivery/http"
	"challenge/internal/repository/memory"
	"challenge/internal/usecase"
)

func main() {
	repo := memory.NewBookMemoryRepository()
	bookUC := usecase.NewBookUsecase(repo)
	authUC := &usecase.AuthUsecase{}

	r := http.NewRouter(bookUC, authUC)
	r.Run(":1234")
}
