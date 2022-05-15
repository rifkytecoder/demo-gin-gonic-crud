package book

import "gorm.io/gorm"

// Responsible for dealing with Databases
type IRepository interface {
	FindAll() ([]Book, error)
	FindById(Id int) (Book, error)
	Create(book Book) (Book, error)
	Update(book Book) (Book, error)
	Delete(book Book) (Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Get list book from table Book
func (r *repository) FindAll() ([]Book, error) {
	var books []Book

	err := r.db.Find(&books).Error

	return books, err
}

// Get book by Id from table Book
func (r *repository) FindById(Id int) (Book, error) {
	var book Book

	err := r.db.Find(&book, Id).Error

	return book, err
}

// Create book into table Book
func (r *repository) Create(book Book) (Book, error) {

	err := r.db.Create(&book).Error

	return book, err
}

// Update book into table Book
func (r *repository) Update(book Book) (Book, error) {

	err := r.db.Save(&book).Error

	return book, err
}

// Delete book into table Book
func (r *repository) Delete(book Book) (Book, error) {

	err := r.db.Delete(&book).Error

	return book, err
}
