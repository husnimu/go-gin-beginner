package initializers

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDatabase() (*gorm.DB, error){
	dsn := "root:@tcp(127.0.0.1:3307)/pustakaapi?charset=utf8mb4&parseTime=True&locLocal"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db,err
}