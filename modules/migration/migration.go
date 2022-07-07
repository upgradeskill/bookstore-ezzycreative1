package migration

import (
	"github.com/upgradeskill/bookstore/modules/book"
	"github.com/upgradeskill/bookstore/modules/user"

	"gorm.io/gorm"
)

func TableMigration(db *gorm.DB) {
	db.AutoMigrate(&user.User{}, &book.Book{})
}
