package memory

import (
	"challenge/internal/domain"
	"errors"
	"sync"
)

type BookMemoryRepository struct {
	data map[string]domain.Book
	mu   sync.RWMutex
}

func NewBookMemoryRepository() *BookMemoryRepository {
	return &BookMemoryRepository{
		data: make(map[string]domain.Book),
	}
}

func (r *BookMemoryRepository) Create(book *domain.Book) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[book.ID] = *book
	return nil
}

func (r *BookMemoryRepository) GetAll() ([]domain.Book, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var books []domain.Book
	for _, v := range r.data {
		books = append(books, v)
	}
	return books, nil
}

func (r *BookMemoryRepository) GetByID(id string) (*domain.Book, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	book, ok := r.data[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return &book, nil
}

func (r *BookMemoryRepository) Update(id string, book *domain.Book) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.data[id]; !ok {
		return errors.New("not found")
	}
	r.data[id] = *book
	return nil
}

func (r *BookMemoryRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.data[id]; !ok {
		return errors.New("not found")
	}
	delete(r.data, id)
	return nil
}

func (r *BookMemoryRepository) DeleteAll() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data = make(map[string]domain.Book)
	return nil
}
