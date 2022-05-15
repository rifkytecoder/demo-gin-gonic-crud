package handlers

import (
	"fmt"
	"lab-gin-crud/book"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Responsible for dealing with EndPoints
type bookHandler struct {
	bookService book.IService
}

func NewBookHandler(bookService book.IService) *bookHandler {
	return &bookHandler{bookService}
}

// Endpoint Get list book
func (h *bookHandler) GetBooks(c *gin.Context) {

	books, err := h.bookService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return

	}

	var booksResponse []book.BookResponse

	for _, b := range books {
		bookResponse := book.BookResponse{
			Id:          b.Id,
			Title:       b.Title,
			Price:       b.Price,
			Description: b.Description,
			Rating:      b.Rating,
			Discount:    b.Discount,
		}

		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

// Endpoint Get book
func (h *bookHandler) GetBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.bookService.FindById(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	bookResponse := book.BookResponse{
		Id:          b.Id,
		Title:       b.Title,
		Price:       b.Price,
		Description: b.Description,
		Rating:      b.Rating,
		Discount:    b.Discount,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

// Endpoint Create book
func (h *bookHandler) CreateBook(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on filed %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})

		return
	}

	b, err := h.bookService.Create(bookRequest)

	bookResponse := book.BookResponse{
		Id:          b.Id,
		Title:       b.Title,
		Price:       b.Price,
		Description: b.Description,
		Rating:      b.Rating,
		Discount:    b.Discount,
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

// Endpoint Get Update book
func (h *bookHandler) UpdateBook(c *gin.Context) {

	var updateBookRequest book.UpdateBookRequest

	err := c.ShouldBindJSON(&updateBookRequest)

	// Validasi Error Messages
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on filed %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusOK, gin.H{
			"error": errorMessages,
		})

		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.bookService.Update(id, updateBookRequest)

	bookResponse := book.BookResponse{
		Id:          b.Id,
		Title:       b.Title,
		Price:       b.Price,
		Description: b.Description,
		Rating:      b.Rating,
		Discount:    b.Discount,
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

// Endpoint Get Delete book
func (h *bookHandler) DeleteBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.bookService.Delete(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := book.BookResponse{
		Id:          b.Id,
		Title:       b.Title,
		Price:       b.Price,
		Description: b.Description,
		Rating:      b.Rating,
		Discount:    b.Discount,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}
