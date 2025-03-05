package handlers

import (
	"backend/internal/app/models"
	"backend/internal/app/servises"
	"backend/pkg/utils"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)


type UserHandler struct {
	serv servises.UserServise
}

func NewUserHandler (serv servises.UserServise) UserHandler{
	return UserHandler{serv: serv}
}

func (h *UserHandler)SignUp(c *gin.Context) {
	var input models.Users
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	input.UserId = utils.GenerateUniqueUserID()
	if input.UserId == "" {
		log.Printf("useridがありません")
	}

	if err := h.serv.ResisterUser(&input); err != nil {
		log.Println("ユーザー登録エラー")
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"missing register User to datadase",
		})
	}

	c.JSON(http.StatusOK, gin.H{"message": "ユーザー作成に成功しました"})
}

func (h *UserHandler) Login(c *gin.Context) {
	var ReqUser struct {
		email    string  
		password string
	}
	if err := c.ShouldBindJSON(&ReqUser);err != nil {
		log.Println(ReqUser)
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Invalid User Request",
		})
	}

	JWTtoken,err := h.serv.ResponseUserIDJWT(ReqUser.email,ReqUser.password)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"missing JWT generation",
		})
	}
	c.SetCookie("jwt",JWTtoken,3600,"/","localhost",false,true)
	c.JSON(http.StatusOK,gin.H{
		"message":"success Generate JWT",
	})
}