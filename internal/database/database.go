package database

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

var err error

func InitDB() {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

		log.Println("database connect write:", dsn)

		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("データベース接続エラー :%v", err)
		}
		log.Println("Success connect DataBase")
	})
}

func GetDB() *gorm.DB {
	return DB
}
