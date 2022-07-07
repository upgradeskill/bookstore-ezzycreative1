package book

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID        string
	Isbn      string
	Title     string
	Price     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ModifyBook struct {
	Isbn      string
	Title     string
	Price     int
	UpdatedAt time.Time
}

type BookSpec struct {
	Isbn  string `validate:"required,len=12"`
	Title string `validate:"required"`
	Price int    `validate:"required"`
}

func toBook(b *BookSpec) *Book {
	uuid := uuid.NewString()
	return &Book{
		ID:        uuid,
		Isbn:      b.Isbn,
		Title:     b.Title,
		Price:     b.Price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func toModifyBook(id string, b *BookSpec) *ModifyBook {
	return &ModifyBook{
		Isbn:      b.Isbn,
		Title:     b.Title,
		Price:     b.Price,
		UpdatedAt: time.Now(),
	}
}
