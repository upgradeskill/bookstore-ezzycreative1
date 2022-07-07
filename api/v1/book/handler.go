package book

import (
	"github.com/upgradeskill/bookstore/api/common"
	"github.com/upgradeskill/bookstore/api/v1/book/request"
	"github.com/upgradeskill/bookstore/api/v1/book/response"
	"github.com/upgradeskill/bookstore/business/book"

	"strconv"

	"github.com/google/uuid"
	echo "github.com/labstack/echo/v4"
)

type BookHandler struct {
	service book.Service
}

func NewBookHandler(service book.Service) *BookHandler {
	return &BookHandler{service}
}

func (c *BookHandler) GetAllBook(ctx echo.Context) error {
	books, err := c.service.GetAllBook()

	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithData(response.AllBook(books)))
}

func (c *BookHandler) FindBookBy(ctx echo.Context) error {
	id := ctx.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	book, err := c.service.FindBookById(id)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithData(response.GetBook(*book)))
}

func (c *BookHandler) AddNewBook(ctx echo.Context) error {
	price, _ := strconv.Atoi(ctx.FormValue("price"))

	book := request.Book{
		Isbn:  ctx.FormValue("isbn"),
		Title: ctx.FormValue("title"),
		Price: price,
	}

	err := c.service.AddNewBook(book.ToBusinessBook())
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}

func (c *BookHandler) ModifyBook(ctx echo.Context) error {
	id := ctx.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	price, _ := strconv.Atoi(ctx.FormValue("price"))

	book := request.Book{
		Isbn:  ctx.FormValue("isbn"),
		Title: ctx.FormValue("title"),
		Price: price,
	}

	err := c.service.ModifyBook(id, book.ToBusinessBook())
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}

func (c *BookHandler) DeleteBook(ctx echo.Context) error {
	id := ctx.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	err := c.service.DeleteBook(id)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}
