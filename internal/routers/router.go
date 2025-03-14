package routers

import (
	"backend/internal/containers"
	"backend/pkg/middleware"
	"backend/pkg/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewAppRouter(app *containers.AppContainer) *gin.Engine {
	router := gin.Default()

	
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},

		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},

		AllowHeaders: []string{
			"Accept",
			"Content-Type",
			"Authorization",
		},

		AllowCredentials: true,
	}))

	user := router.Group("/user")
	api := router.Group("/api")
	
	user.POST("/signup",app.UserHandler.SignUp)
	user.POST("/login",app.UserHandler.Login)
	user.POST("/logout",app.UserHandler.Logout)
	user.GET("/ckeckauth",utils.CkeckAuth)

	api.GET("/userdetail",app.UserHandler.ResponseUserDetail).Use(middleware.AuthMiddleware())
	

	classes := router.Group("/classes")
	classes.Use(middleware.AuthMiddleware())
	
	
	classes.POST("/register",app.ClassesHandler.RegisterClass)
	classes.POST("/delete",app.ClassesHandler.DeleteRegisteredClass)
	classes.POST("/classes",app.ClassesHandler.ViewUserClasses)
	classes.GET("/classes",app.ClassesHandler.ViewUserClassesByUserID)
	classes.GET("/schedule",app.ClassesHandler.ViewUserSchedule)
	classes.GET("/checktool",app.ClassesHandler.CheckToolAPI)
	
	classes.GET("/userdetail",app.UserHandler.ResponseUserDetail)
	
	return router 
}