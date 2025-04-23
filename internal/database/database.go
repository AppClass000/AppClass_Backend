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
				
		    user := os.Getenv("MYSQLUSER")
			password := os.Getenv("MYSQLPASSWORD")
		    host := os.Getenv("MYSQLHOST")
			port := os.Getenv("MYSQLPORT")
			dbname := os.Getenv("MYSQLDATABASE")

			dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",user, password, host, port, dbname)
			
		

		log.Println("database connect write:  %s", dsn)

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
