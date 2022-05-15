package book

import "encoding/json"

// Request Body
type BookRequest struct {
	Title       string      `json:"title" binding:"required"`
	Price       json.Number `json:"price" binding:"required"`
	Description string      `json:"description" binding:"required"`
	Rating      json.Number `json:"rating" binding:"required"`
	Discount    json.Number `json:"discount" binding:"required"`
}
type UpdateBookRequest struct {
	Title       string      `json:"title" binding:"required"`
	Price       json.Number `json:"price" binding:"required"`
	Description string      `json:"description"`
	Rating      json.Number `json:"rating"`
	Discount    json.Number `json:"discount"`
}
