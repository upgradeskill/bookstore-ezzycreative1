package book

import (
	"github.com/upgradeskill/bookstore/business"
	"github.com/upgradeskill/bookstore/util/validator"
)

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) GetAllBook() (*[]Book, error) {
	return s.repository.GetAllBook()
}

func (r *service) FindBookById(id string) (*Book, error) {
	return r.repository.FindBookById(id)
}

func (s *service) AddNewBook(book *BookSpec) error {
	err := validator.GetValidator().Struct(book)
	if err != nil {
		return business.ErrInvalidSpec
	}
	return s.repository.AddNewBook(toBook(book))
}

func (s *service) ModifyBook(id string, book *BookSpec) error {
	err := validator.GetValidator().Struct(book)
	if err != nil {
		return business.ErrInvalidSpec
	}

	return s.repository.ModifyBook(id, toModifyBook(id, book))
}

func (s *service) DeleteBook(id string) error {
	return s.repository.DeleteBook(id)
}
