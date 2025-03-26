package utils

import (
	"backend/internal/app/repositories"
	"backend/internal/database"
	"backend/pkg/auth"
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
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
	ok := auth.VaridateJWT(tokenString)
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

func VaridateUserID (c *gin.Context) (string,error){
	value, exist := c.Get("userID")
	if !exist {
		log.Println("UserIDがありません:ResponseUserDeail")
	}
	userId, ok := value.(string)
	if !ok {
		log.Println("Invalid UserID Type")
		return "",fmt.Errorf("error varidate userid")
	}

	return userId ,nil
}
