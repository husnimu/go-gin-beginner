package initializers

import (
	"os/user"
	"pustaka-api/book"

	"gorm.io/gorm"
)

func SyncDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(book.Book{}, user.User{})
	return err
}