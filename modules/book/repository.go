package book

import (
	"errors"
	"time"

	"github.com/upgradeskill/bookstore/business"
	"github.com/upgradeskill/bookstore/business/book"

	"gorm.io/gorm"
)

type Book struct {
	ID        string    `gorm:"id;type:uuid;primaryKey"`
	Isbn      string    `gorm:"code;type:varchar(12);index:book_code_uniq;unique"`
	Title     string    `gorm:"title;type:varchar(100)"`
	Price     int       `gorm:"price;default:0"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
	DeletedAt time.Time `gorm:"deleted_at;default:'0001-01-01 00:00:00+00'"`
}

func (b *Book) toBusinessBook() *book.Book {
	return &book.Book{
		ID:        b.ID,
		Isbn:      b.Isbn,
		Title:     b.Title,
		Price:     b.Price,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}
}

func insertBook(b *book.Book) *Book {
	return &Book{
		ID:        b.ID,
		Isbn:      b.Isbn,
		Title:     b.Title,
		Price:     b.Price,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}
}

func allBusinessBook(books *[]Book) *[]book.Book {
	var items []book.Book

	for _, item := range *books {
		items = append(items, *item.toBusinessBook())
	}

	if items == nil {
		items = []book.Book{}
	}

	return &items
}

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetAllBook() (*[]book.Book, error) {
	books := new([]Book)

	err := r.DB.Where("(deleted_at IS NULL OR deleted_at = '0001-01-01 00:00:00+00')").Find(books).Error
	if err != nil {
		return nil, err
	}

	return allBusinessBook(books), nil
}

func (r *Repository) FindBookById(id string) (*book.Book, error) {
	item := new(Book)

	err := r.DB.First(item, "ID = ? and (deleted_at IS NULL OR deleted_at = '0001-01-01 00:00:00+00')", id).Error
	if err != nil {
		return nil, err
	}

	return item.toBusinessBook(), nil
}

func (r *Repository) AddNewBook(book *book.Book) error {
	var err error

	err = r.DB.First(&Book{}, " code = ?", book.Isbn).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return business.ErrConflig
	}

	item := insertBook(book)

	err = r.DB.Create(item).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) ModifyBook(id string, book *book.ModifyBook) error {
	var err error
	var item Book

	err = r.DB.First(&item, "ID = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return business.ErrNotFound
		}
		return err
	}

	var tmp Book
	err = r.DB.First(&tmp, " isbn = ? ", book.Isbn).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		if err == nil {
			if tmp.ID != id {
				return business.ErrConflig
			}
		}
	}

	return r.DB.Model(&item).Updates(Book{
		Isbn:      book.Isbn,
		Title:     book.Title,
		Price:     book.Price,
		UpdatedAt: book.UpdatedAt,
	}).Error
}

func (r *Repository) DeleteBook(id string) error {
	book := new(Book)

	err := r.DB.First(book, "ID = ? and (deleted_at IS NULL OR deleted_at = '0001-01-01 00:00:00+00')", id).Error
	if err != nil {
		return err
	}

	return r.DB.Model(book).Updates(Book{DeletedAt: time.Now()}).Error
}
