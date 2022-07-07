package request

import "github.com/upgradeskill/bookstore/business/book"

type Book struct {
	Isbn   string `json:"isbn"`
	Title  string `json:"name"`
	Author string `json:"author"`
	Price  int    `json:"price"`
}

func (b *Book) ToBusinessBook() *book.BookSpec {
	return &book.BookSpec{
		Isbn:  b.Isbn,
		Title: b.Title,
		Price: b.Price,
	}
}
