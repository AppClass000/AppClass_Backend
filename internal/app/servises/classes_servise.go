package controllers

import (
	"backend/database"
	"backend/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


func SignUp(c *gin.Context) {
	var input models.Users
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if input.Name == "" || input.Email == "" {
		log.Printf("名前がからです ")
	}

	input.UserId = utils.GenerateUniqueUserID()
	if input.UserId == "" {
		log.Printf("useridがありません")
	}
	log.Println(input.UserId, input.Name, input)

	if err := models.CreateUser(database.DB, &input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "ユーザー作成失敗"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ユーザー作成に成功しました"})
}

func ResponseJWT(c *gin.Context) {
	var reqUser models.ReqUsers
	var user models.Users

	c.SetSameSite(http.SameSiteLaxMode)

	if err := c.ShouldBindJSON(&reqUser); err != nil {
		log.Printf("userをバインド出来ませんでした: %v", err)
	}
	log.Println("email:", reqUser.Email)
	query := database.DB.Where("Email = ?", reqUser.Email)
	if err := query.Select("UserId").First(&user).Error; err != nil {
		log.Println("usersバインドエラー:", err)
	}
	log.Println("渡されるuserid:", user.UserId)
	token, err := utils.GenerateJWT(user.UserId)
	if err != nil {
		log.Printf("JWT生成でエラーが発生しました: %v", err)
	}
	if token == "" {
		log.Printf("tokenが空です")
	}

	c.SetCookie("jwt", token, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "jWT ok",
	})

}

func ResponseUserClasses(c *gin.Context) {
	var user_detail models.UserDetail

	user_id, exist := c.Get("userID")
	if !exist {
		log.Printf("userIDが存在しません")
		return
	}
	query := database.DB.Select("faculty", "department", "course").Where("user_id = ?", user_id)
	err := query.Find(&user_detail).Error
	if err != nil {
		log.Println("user_detailのバインドでエラーがおきました", err)
	}
	filter := user_detail
	classes, err := models.GetClasses(&filter)
	if err != nil {
		log.Println("授業取得に失敗しました:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"classes": classes,
	})
}

func GetScheduleData(c *gin.Context) {
	UserID, exists := c.Get("userID")
	if !exists {
		log.Println("userIDが存在しません")
		return
	}
	UserId, ok := UserID.(string)
	if !ok {
		log.Println("UserIDのstring型への変換に失敗しました")
		return
	}

	UserClasses, err := models.GetUserClasses(UserId)
	if err != nil {
		log.Println("userの登録データ取得エラー :%V", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":      "success",
		"user_classes": UserClasses,
	})

}

func CheckUserResitered(c *gin.Context) {
 	var IsCoreList []int
 	var IsIntroductoryList []int
 	var IsCommonList []int


 	UserID, ok := c.Get("userID")
 	if !ok {
 		log.Panic("userID don't exist")
 	}
	UserId, ok := UserID.(string)
 	if !ok {
 		log.Println("mistake convert UserID to string")
// 	}
// 	userClasses, err := models.GetUserClasses(UserId)
// 	if err != nil {
// 		log.Println("Getting userCLasses error")
// 	}

// 	for i :=0; i< len(userClasses); i++ {
// 		if userClasses[i].IsCore == 1 {
// 			IsCoreList := append(IsCoreList,userClasses[i].IsCore)
// 		}
// 	    if userClasses[i].IsIntroductory == 1 {
// 			IsIntroductoryList := append(IsIntroductoryList,userClasses[i].IsIntroductory)
// 		}
// 	    if userClasses[i].IsCommon == 1 {
// 	    	IsCommonList := append(IsCommonList,userClasses[i].IsCommon)
// 	    }
// 	}

// 	userDetail,err := models.GetUserDetail(database.DB,UserId)
// 	if err != nil {
// 		log.Println("error of GetUserDetail:",err)
// 	}
// 	if userDetail.Faculty == "工学部" {
		
// 	}




// }
// // 