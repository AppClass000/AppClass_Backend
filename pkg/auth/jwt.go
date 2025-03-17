package auth


import (
	"backend/config"
	
	"fmt"
	"log"
	"time"
	"github.com/golang-jwt/jwt/v5"

)


func GenerateJWT(userID string) (string, error) {

	config.LoadEnv()

	log.Println("これがuserid", userID)
	var jwtSecretKey = config.GetEnv("JWTSECRETKEY","")
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

func VaridateJWT(tokenString string) bool {
	config.LoadEnv()

	token, err := jwt.Parse(tokenString,func(t *jwt.Token) (interface{}, error) {
		return []byte(config.GetEnv("JWTSECRETKEY","")),nil
	})
	if err != nil || !token.Valid{
		log.Println("token parse error:",err)
		return false
	}
	return true
}
