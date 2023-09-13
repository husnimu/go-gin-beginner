package initializers

import (
	"pustaka-api/book"
	"pustaka-api/user"

	"gorm.io/gorm"
)

func SyncDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(user.User{}, book.Book{})
	return err
}
