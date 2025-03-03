package controllers

import (
	"backend/database"
	"backend/models"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var input models.Users

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := hashPassword(input.Password)
	if err != nil {
		log.Fatalf("パスワード暗号化error: %v", err)
	}

	newUser := models.Users{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "データベースに登録できませんでした"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "登録されました",
		"user":    newUser,
	})
}

func hashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func RegisterUserClasses(c *gin.Context) {
	var userClass models.UserClasses
	err := c.ShouldBindBodyWithJSON(&userClass)
	if err != nil {
		log.Printf("userClassのbindに失敗しました: %v", err)
	}
	userID, exist := c.Get("userID")
	if !exist {
		log.Println("userIDがありません")
	}
	userClass.UserId = userID.(string)

	err = models.CreateClasses(&userClass)
	if err != nil {
		return
	}
	log.Println("授業が登録されました")
	c.JSON(http.StatusOK, gin.H{
		"message": "Registered!",
	})
}

func ResponseClasses(c *gin.Context) {
	var filter models.UserDetail

	err := c.ShouldBindJSON(&filter)
	if err != nil {
		log.Println("filterのバインドエラー", err)
	}
	classes, err := models.GetClasses(&filter)
	if err != nil {
		log.Println("授業取得エラー", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "授業取得に失敗",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "授業取得成功",
		"classes": classes,
	})
}
