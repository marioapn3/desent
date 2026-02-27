package usecase

import (
	"challenge/internal/domain"
	"challenge/internal/repository"
	"errors"
	"strconv"
	"sync/atomic"
)

var idCounter atomic.Int64

type BookUsecase struct {
	repo repository.BookRepository
}

func NewBookUsecase(r repository.BookRepository) *BookUsecase {
	return &BookUsecase{r}
}

func (u *BookUsecase) Create(book *domain.Book) error {
	if book.Title == "" || book.Author == "" {
		return errors.New("invalid input")
	}
	if book.ID == "" {
		book.ID = strconv.FormatInt(idCounter.Add(1), 10)
	}
	return u.repo.Create(book)
}

func (u *BookUsecase) GetAll(author string, page, limit int) ([]domain.Book, error) {
	books, _ := u.repo.GetAll()

	var filtered []domain.Book
	for _, b := range books {
		if author == "" || b.Author == author {
			filtered = append(filtered, b)
		}
	}

	if limit <= 0 {
		return filtered, nil
	}

	start := (page - 1) * limit
	end := start + limit

	if start >= len(filtered) {
		return []domain.Book{}, nil
	}
	if end > len(filtered) {
		end = len(filtered)
	}
	return filtered[start:end], nil
}

func (u *BookUsecase) GetByID(id string) (*domain.Book, error) {
	return u.repo.GetByID(id)
}

func (u *BookUsecase) Update(id string, book *domain.Book) error {
	return u.repo.Update(id, book)
}

func (u *BookUsecase) Delete(id string) error {
	return u.repo.Delete(id)
}

func (u *BookUsecase) DeleteAll() error {
	return u.repo.DeleteAll()
}
