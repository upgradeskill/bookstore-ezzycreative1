package book

type Service interface {
	// Get all book
	GetAllBook() (*[]Book, error)

	// Get book detail by id
	FindBookById(id string) (*Book, error)

	// Add new book
	AddNewBook(book *Book) error

	// Modify Information book
	ModifyBook(id string, bokk *BookSpec) error

	// Delete book
	DeleteBook(id string) error
}

type Repository interface {
	// Get all book
	GetAllBook() (*[]Book, error)

	// Get book detail by id
	FindBookById(id string) (*Book, error)

	// Add new book
	AddNewBook(book *Book) error

	// Modify Information Book
	ModifyBook(id string, book *ModifyBook) error

	// Delete book
	DeleteBook(id string) error
}
