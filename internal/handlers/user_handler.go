package handlers

import (
	"backend/internal/app/models"
	"backend/internal/app/services"
	"backend/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	serv services.UserServise
}

func NewUserHandler(serv services.UserServise) UserHandler {
	return UserHandler{serv: serv}
}

func (h *UserHandler) SignUp(c *gin.Context) {
	var input models.Users
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	input.UserID = utils.GenerateUniqueUserID()
	if input.UserID == "" {
		log.Printf("useridがありません")
		return
	}

	if err := h.serv.ResisterUser(&input); err != nil {
		log.Println("ユーザー登録エラー")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "missing register User to datadase",
		})
		return
	}

	JWTtoken, err := h.serv.ResponseSignUpJTW(input.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "missing JWT generation",
		})
		return
	}
	c.SetCookie("jwt", JWTtoken, 3600, "/", "appclass.up.railway.app", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "success Generate JWT and registered User",
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	var ReqUser struct {
		Email    string `json:"email"`
		Password string `json':"password"`
	}

	if err := c.ShouldBindJSON(&ReqUser); err != nil {
		log.Println(ReqUser)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid User Request",
		})
		return
	}

	JWTtoken, err := h.serv.ResponseUserIDJWT(ReqUser.Email, ReqUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "missing JWT generation",
		})
		return
	}
	c.SetCookie("jwt", JWTtoken, 3600, "/", "appclass.up.railway.app", true, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "success Generate JWT",
	})
}

func (h *UserHandler) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", 0, "/", "appclass.up.railway.app", true, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "logout success",
	})
}

func (h *UserHandler) RegisterUserDetail(c *gin.Context) {
	var userdetail models.UserDetail

	userID, err := utils.VaridateUserID(c)
	if err != nil {
		log.Println("Invalid UserID Type")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid userid",
		})
		return
	}

	err = c.ShouldBindJSON(&userdetail)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid userdetail",
		})
		return
	}
	err = h.serv.RegisterUserDetail(&userdetail, userID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid userdetail",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})

}

func (h *UserHandler) ResponseUserDetail(c *gin.Context) {
	value, exist := c.Get("userID")
	if !exist {
		log.Println("UserIDがありません:ResponseUserDeail")
	}
	userId, ok := value.(string)
	if !ok {
		log.Println("Invalid UserID Type")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid userid",
		})
		return
	}
	userDetail, err := utils.GetUserDetail(userId)
	if err != nil {
		log.Println("missing get userDetail by userid:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "missing get userDetail by userid",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "success",
		"userdetail": userDetail,
	})

}

func (h *UserHandler) RegisterUserNameHandle(c *gin.Context) {
	var userName models.RequestUserName

	if err := c.ShouldBindJSON(&userName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json",
		})
		return
	}
	userID, err := utils.VaridateUserID(c)
	if err != nil {
		log.Println("err:", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	user, err := h.serv.GetUserByUserID(userID)
	if err != nil {
		log.Println("err:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	err = h.serv.ResisterUserName(user, userName.Name)
	if err != nil {
		log.Println("err:", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})

}

func (h *UserHandler) ResponseUserProfile(c *gin.Context) {

	userID, err := utils.VaridateUserID(c)
	if err != nil {
		log.Println("err:", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	user, err := h.serv.GetUserByUserID(userID)
	if err != nil {
		log.Println("err:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK,
		gin.H{"name": user.Name, "email": user.Email},
	)

}
