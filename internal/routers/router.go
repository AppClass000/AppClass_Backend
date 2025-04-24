package routers

import (
	"backend/internal/containers"
	"backend/pkg/auth"
	"backend/pkg/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewAppRouter(app *containers.AppContainer) *gin.Engine {
	router := gin.Default()
	

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"https://appclass.up.railway.app",
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

	authUser := user.Group("/")
	authUser.Use(auth.AuthMiddleware())

	user.POST("/signup", app.UserHandler.SignUp)
	user.POST("/login", app.UserHandler.Login)
	user.POST("/logout", app.UserHandler.Logout)
	user.GET("/ckeckauth", utils.CkeckAuth)

	authUser.GET("/profile", app.UserHandler.ResponseUserProfile)
	authUser.GET("/userdetail", app.UserHandler.ResponseUserDetail)
	authUser.POST("/userdetail", app.UserHandler.RegisterUserDetail)
	authUser.POST("/username", app.UserHandler.RegisterUserNameHandle)

	classes := router.Group("/classes")
	classes.Use(auth.AuthMiddleware())

	classes.POST("/register", app.ClassesHandler.RegisterClass)
	classes.POST("/delete", app.ClassesHandler.DeleteRegisteredClass)
	classes.POST("/classes", app.ClassesHandler.ViewUserClasses)
	classes.GET("/classes", app.ClassesHandler.ViewUserClassesByUserID)
	classes.GET("/schedule", app.ClassesHandler.ViewUserSchedule)
	classes.GET("/checktool", app.ClassesHandler.CheckToolAPI)

	return router
}
