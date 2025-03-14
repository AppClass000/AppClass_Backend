package utils

import (
	"backend/config"
	"backend/internal/app/repositories"
	"backend/internal/database"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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


func GetUserDetail(userID string) (*repositories.UserDetail,error)  {
	var userDetail repositories.UserDetail
	db := database.GetDB()

	query := db.Select("faculty","department","course")

	if err :=query.Where("user_id = ?",userID).First(&userDetail).Error;err != nil {
		return nil,err 
	}
	log.Println("Executing SQL",query.Statement.SQL.String())
	return &userDetail,nil 
}

func CkeckAuth(c *gin.Context) {
	tokenString,err := c.Cookie("jwt")
	if  tokenString == "" {
		c.JSON(http.StatusBadRequest,gin.H{
			"message":"unauthorized",
		})
		log.Println(tokenString,"cokkie debug log:",err)
		return
	}
	ok := VaridateJWT(tokenString)
	if   !ok {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"JWT Invalid",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"authorized",
	})
}