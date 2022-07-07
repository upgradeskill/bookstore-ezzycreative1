package api

import (
	"github.com/upgradeskill/bookstore/api/middleware"
	"github.com/upgradeskill/bookstore/api/v1/book"
	"github.com/upgradeskill/bookstore/api/v1/user"

	echo "github.com/labstack/echo/v4"
)

func AddRoute(e *echo.Echo, user *user.UserHandler, book *book.BookHandler) {
	if user == nil || book == nil {
		panic("Invalid route parameters")
	}

	e.POST("/v1/register", user.AddNewUser)
	e.POST("/v1/login", user.LoginUser)

	eproduct := e.Group("/v1/products")
	eproduct.Use(middleware.JWTMiddleware())
	eproduct.GET("", book.GetAllBook)
	eproduct.POST("", book.AddNewBook)
	eproduct.GET("/:id", book.FindBookBy)
	eproduct.PUT("/:id", book.ModifyBook)
	eproduct.DELETE("/:id", book.DeleteBook)
}
