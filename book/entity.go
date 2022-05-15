package book

import "time"

// Table Book
type Book struct {
	Id          int
	Title       string
	Description string
	Price       int
	Rating      int
	Discount    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
