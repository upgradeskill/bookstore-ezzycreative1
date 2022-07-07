package main

import (
	"github.com/upgradeskill/bookstore/api"
	bookHandler "github.com/upgradeskill/bookstore/api/v1/book"
	userController "github.com/upgradeskill/bookstore/api/v1/user"

	bookService "github.com/upgradeskill/bookstore/business/book"
	userService "github.com/upgradeskill/bookstore/business/user"

	"github.com/upgradeskill/bookstore/config"
	"github.com/upgradeskill/bookstore/modules/migration"

	bookRepository "github.com/upgradeskill/bookstore/modules/book"
	userRepository "github.com/upgradeskill/bookstore/modules/user"

	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	echo "github.com/labstack/echo/v4"
)

func newDatabaseConnection(config *config.ConfigApp) *gorm.DB {
	stringConfig := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DbUsername,
		config.DbPassword,
		config.DbHost,
		config.DbPort,
		config.DbName,
	)

	db, err := gorm.Open(mysql.Open(stringConfig), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	migration.TableMigration(db)

	return db
}

func main() {
	// Geting config app
	config := config.GetConfig()

	// Initiate database connection
	dbConnection := newDatabaseConnection(config)

	// Initiate User Repository
	userRepo := userRepository.NewRepository(dbConnection)

	// Initiate User Service
	userServc := userService.NewService(userRepo)

	// Initiate User Controller
	userHandler := userController.NewUserHandler(userServc)

	// Initiate Book Repository
	bookRepo := bookRepository.NewRepository(dbConnection)

	// Initiate Book Service
	bookServc := bookService.NewService(bookRepo)

	// Initiate Product Controller
	bookHandler := bookHandler.NewBookHandler(bookServc)

	// Initiate Echo Web Service
	e := echo.New()

	// Add routing
	api.AddRoute(e, userHandler, bookHandler)

	// Start service
	e.Start(fmt.Sprintf("%s:%d", config.AppHost, config.AppPort))
}
