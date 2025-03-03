package db

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

type Users struct {
	Id        uint `gorm:"primarykey"`
	UserId    string
	Name      string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("環境変数file読み込みエラー :%v", err)
	}
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatalf("DSNが設定されていません")
	}

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("データベース接続エラー :%v", err)
	}

}
