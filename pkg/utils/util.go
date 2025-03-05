package utils

import (
	"fmt"
	"log"
	"os"
	"time"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func GenerateUniqueUserID() string {
	return uuid.NewString()
}

func IsUserLoggedIn(c *gin.Context) (string, bool) {
	session := sessions.Default(c)
	userID, ok := session.Get("user_id").(string)
	return userID, ok
}

func GenerateJWT(userID string) (string, error) {
	envPath := "C:/project/AppClass/backend/.env"
	err := godotenv.Load(envPath)
	if err != nil {
		log.Println(".envファイルを読み込めませんでした:", err)
	}

	log.Println("これがuserid", userID)
	var jwtSecretKey = os.Getenv("JWTSECRETKEY")
	log.Println("これが秘密鍵", jwtSecretKey)
	if jwtSecretKey == "" {
		log.Fatalf("秘密鍵が取得できませんでした")
	}

	claim := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 20).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString([]byte(jwtSecretKey))

	if err != nil {
		return "", fmt.Errorf("JWT署名時エラーが発生しました:%w", err)
	}
	return tokenString, nil
}

func CreateJWTResponse(userID string) (bool, string) {
	var loggedIn bool
	tokenString, err := GenerateJWT(userID)
	if err != nil {
		log.Printf("トークン作成時にエラーが発生しました: %v", err)
		loggedIn = false
		return loggedIn, ""
	}
	loggedIn = true

	return loggedIn, tokenString

}
