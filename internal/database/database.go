package database

import (
	"backend/config"
	"gorm.io/driver/mysql"
	"sync"
	"gorm.io/gorm"
	"log"
)

var (DB *gorm.DB
	once sync.Once
)

var err error
func InitDB() {
	once.Do(func() {config.LoadEnv()
	
	dsn := config.GetEnv("DATABASE_URL","default")
	if dsn == "default" {
		log.Fatalf("環境変数が取得できない")
	}
	
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("データベース接続エラー :%v", err)
	}
	log.Println("Success connect DataBase")
    })
}

func GetDB () *gorm.DB{
	return DB
}
