package book

// Responsible for dealing with Business logic
type IService interface {
	FindAll() ([]Book, error)
	FindById(Id int) (Book, error)
	Create(BookRequest BookRequest) (Book, error)
	Update(id int, updateBookRequest UpdateBookRequest) (Book, error)
	Delete(id int) (Book, error)
}

type service struct {
	repository IRepository
}

func NewService(repository IRepository) *service {
	return &service{repository}
}

// Get list book
func (s *service) FindAll() ([]Book, error) {

	books, err := s.repository.FindAll()

	return books, err
}

// Get book by Id
func (s *service) FindById(id int) (Book, error) {

	book, err := s.repository.FindById(id)

	return book, err
}

// Create book
func (s *service) Create(bookRequest BookRequest) (Book, error) {

	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	discount, _ := bookRequest.Discount.Int64()

	book := Book{
		Title:       bookRequest.Title,
		Price:       int(price),
		Description: bookRequest.Description,
		Rating:      int(rating),
		Discount:    int(discount),
	}

	newBook, err := s.repository.Create(book)

	return newBook, err
}

// Update book
func (s *service) Update(id int, updateBookRequest UpdateBookRequest) (Book, error) {

	book, _ := s.repository.FindById(id)

	price, _ := updateBookRequest.Price.Int64()
	rating, _ := updateBookRequest.Rating.Int64()
	discount, _ := updateBookRequest.Discount.Int64()

	book.Title = updateBookRequest.Title
	book.Price = int(price)
	book.Description = updateBookRequest.Description
	book.Rating = int(rating)
	book.Discount = int(discount)

	newBook, err := s.repository.Update(book)
	return newBook, err
}

// Delete book
func (s *service) Delete(id int) (Book, error) {

	book, _ := s.repository.FindById(id)

	newBook, err := s.repository.Delete(book)

	return newBook, err
}
