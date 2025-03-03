package router

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func SignUproute(r *gin.Engine) {
	r.POST("/signup", controllers.SignUp)
}

func UserDetailroute(r *gin.RouterGroup) {
	r.POST("/userdetail", controllers.ResponseClasses)
}

func ResponseJWTroute(r *gin.Engine) {
	r.POST("/login", controllers.ResponseJWT)
}

func RegisterClassesroute(r *gin.RouterGroup) {
	r.POST("/classesregister", controllers.RegisterUserClasses)
}

func GetUserClasssesroute(r *gin.RouterGroup) {
	r.GET("/userclasses", controllers.ResponseUserClasses)
}

func GetScheduleDataroute(r *gin.RouterGroup) {
	r.GET("/schedule", controllers.GetScheduleData)
}

func CheckToolAPI(r *gin.RouterGroup) {
	r.POST("/checktool")
}
