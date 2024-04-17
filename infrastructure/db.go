package infrastructure

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/test-invoice?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("db init error: ", err)
	}
	fmt.Println("success db connction")
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
