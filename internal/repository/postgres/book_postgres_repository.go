package postgres

import (
	"challenge/internal/domain"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BookPostgresRepository struct {
	pool *pgxpool.Pool
}

func NewBookPostgresRepository(pool *pgxpool.Pool) *BookPostgresRepository {
	return &BookPostgresRepository{pool: pool}
}

func (r *BookPostgresRepository) InitTable() error {
	_, err := r.pool.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS books (
			id TEXT PRIMARY KEY,
			title TEXT NOT NULL,
			author TEXT NOT NULL,
			year INTEGER NOT NULL DEFAULT 0
		)
	`)
	return err
}

func (r *BookPostgresRepository) Create(book *domain.Book) error {
	_, err := r.pool.Exec(context.Background(),
		"INSERT INTO books (id, title, author, year) VALUES ($1, $2, $3, $4)",
		book.ID, book.Title, book.Author, book.Year,
	)
	return err
}

func (r *BookPostgresRepository) GetAll() ([]domain.Book, error) {
	rows, err := r.pool.Query(context.Background(), "SELECT id, title, author, year FROM books ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []domain.Book
	for rows.Next() {
		var b domain.Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Year); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

func (r *BookPostgresRepository) GetByID(id string) (*domain.Book, error) {
	var b domain.Book
	err := r.pool.QueryRow(context.Background(),
		"SELECT id, title, author, year FROM books WHERE id = $1", id,
	).Scan(&b.ID, &b.Title, &b.Author, &b.Year)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("not found")
		}
		return nil, err
	}
	return &b, nil
}

func (r *BookPostgresRepository) Update(id string, book *domain.Book) error {
	result, err := r.pool.Exec(context.Background(),
		"UPDATE books SET title = $1, author = $2, year = $3 WHERE id = $4",
		book.Title, book.Author, book.Year, id,
	)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return errors.New("not found")
	}
	return nil
}

func (r *BookPostgresRepository) Delete(id string) error {
	result, err := r.pool.Exec(context.Background(),
		"DELETE FROM books WHERE id = $1", id,
	)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return errors.New("not found")
	}
	return nil
}

func (r *BookPostgresRepository) DeleteAll() error {
	_, err := r.pool.Exec(context.Background(), "DELETE FROM books")
	return err
}
