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
		    host := os.Getenv("MYSQLHOST")
			port := os.Getenv("MYSQLPORT")
			user := os.Getenv("MYSQLUSER")
			password := os.Getenv("MYSQLPASSWORD")
			dbname := os.Getenv("MYSQLDATABASE")
			
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",host,port,user,password,dbname,
		)
		
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
