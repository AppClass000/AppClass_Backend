package utils

import (
	"testing"

	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func TestGenerateJWT(t *testing.T) {
	envPath := "/home/level/project/AppClass/backend/config/.env"
	err := godotenv.Load(envPath)
	if err != nil {
		log.Println(".envファイルを読み込めませんでした:", err)
	}
	secretkey := os.Getenv("JWTSECRETKEY")
	log.Println("検証用の鍵:", secretkey)

	user_id := "222"

	tokenString, err := GenerateJWT(user_id)
	if err != nil {
		t.Fatalf("GenerateJWTでエラー発生 %v", err)
	}

	log.Println("生成されたtoken:", tokenString)
	parseToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretkey), nil
	})

	if err != nil || parseToken.Valid {
		log.Printf("生成されたjwtが無効です %v", err)
	}

	claims, ok := parseToken.Claims.(jwt.MapClaims)
	if !ok {
		t.Fatalf("クレームの取得に失敗しました %v", err)
	}

	if ok && claims["exp"] != nil {
		log.Println("Token Expiration:", time.Unix(int64(claims["exp"].(float64)), 0))
	}

	if claims["user_id"] != user_id {
		t.Errorf("期待されるuser_idと違います: got %v want %v", claims["user_id"], user_id)
	}

}
