package config 

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)


func LoadEnv () {
	err := godotenv.Load("/home/level/project/AppClass/backend/config/.env")
	if err != nil {
		log.Fatalf("環境変数の読み込みに失敗しました %v",err)
	}
	
}

func GetEnv (key string,defaultvalue string) string{
	value,exist := os.LookupEnv(key)
	if !exist {
		log.Println("環境変数 %v が存在しません",key)
		return defaultvalue
	}

	return value
}