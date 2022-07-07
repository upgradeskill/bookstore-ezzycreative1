package response

import (
	"time"

	"github.com/upgradeskill/bookstore/business/book"
)

type DetailBook struct {
	ID        string
	Isbn      string
	Title     string
	Price     int
	UpdatedAt time.Time
}

func GetBook(b book.Book) DetailBook {
	return DetailBook{
		ID:        b.ID,
		Isbn:      b.Isbn,
		Title:     b.Title,
		Price:     b.Price,
		UpdatedAt: b.UpdatedAt,
	}
}

func AllBook(books *[]book.Book) *[]DetailBook {
	var listBook []DetailBook

	for _, item := range *books {
		listBook = append(listBook, GetBook(item))
	}

	if listBook == nil {
		listBook = []DetailBook{}
	}

	return &listBook
}
