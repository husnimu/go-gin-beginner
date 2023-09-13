package book

type Service interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
	Update(ID int, updateBookRequest UpdateBookRequest) (Book, error)
	Delete(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
}

func (s *service) FindById(ID int) (Book, error) {
	book, err := s.repository.FindById(ID)
	return book, err
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	book := Book{
		Title:       bookRequest.Title,
		Price:       bookRequest.Price,
		Description: bookRequest.Description,
		Rating:      bookRequest.Rating,
	}

	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) Update(ID int, updateBookRequest UpdateBookRequest) (Book, error) {
	book, err := s.repository.FindById(ID)
	if err != nil {
		return book, err
	}

	if updateBookRequest.Title != "" {
		book.Title = updateBookRequest.Title
	}
	if updateBookRequest.Price != 0 {
		book.Price = updateBookRequest.Price
	}
	if updateBookRequest.Description != "" {
		book.Description = updateBookRequest.Description
	}
	if updateBookRequest.Rating != 0 {
		book.Rating = updateBookRequest.Rating
	}

	newBook, err := s.repository.Update(book)
	return newBook, err
}

func (s *service) Delete(ID int) (Book, error) {
	book, err := s.repository.FindById(ID)
	_, err = s.repository.Delete(book)

	return book, err
}
