package handler

import (
	"challenge/internal/domain"
	"challenge/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	usecase *usecase.BookUsecase
}

func NewBookHandler(u *usecase.BookUsecase) *BookHandler {
	return &BookHandler{u}
}

func (h *BookHandler) Create(c *gin.Context) {
	var book domain.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid"})
		return
	}
	err := h.usecase.Create(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, book)
}

func (h *BookHandler) GetAll(c *gin.Context) {
	author := c.Query("author")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	if page == 0 {
		page = 1
	}

	books, _ := h.usecase.GetAll(author, page, limit)
	if books == nil {
		books = []domain.Book{}
	}
	c.JSON(http.StatusOK, books)
}

func (h *BookHandler) GetByID(c *gin.Context) {
	book, err := h.usecase.GetByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) Update(c *gin.Context) {
	id := c.Param("id")
	existing, err := h.usecase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	var book domain.Book
	c.ShouldBindJSON(&book)
	book.ID = id
	if book.Title == "" {
		book.Title = existing.Title
	}
	if book.Author == "" {
		book.Author = existing.Author
	}
	if book.Year == 0 {
		book.Year = existing.Year
	}
	h.usecase.Update(id, &book)
	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) Delete(c *gin.Context) {
	err := h.usecase.Delete(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
