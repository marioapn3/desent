package main

import (
	"context"
	"log"
	"os"

	"challenge/internal/delivery/http"
	"challenge/internal/repository/postgres"
	"challenge/internal/usecase"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbURL := os.Getenv("DATABASE_URL")
	log.Println("DATABASE_URL:", dbURL)
	if dbURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	defer pool.Close()

	repo := postgres.NewBookPostgresRepository(pool)
	if err := repo.InitTable(); err != nil {
		log.Fatal("failed to init table:", err)
	}

	bookUC := usecase.NewBookUsecase(repo)
	authUC := &usecase.AuthUsecase{}

	port := os.Getenv("PORT")
	if port == "" {
		port = "1234"
	}

	r := http.NewRouter(bookUC, authUC)
	r.Run(":" + port)
}
