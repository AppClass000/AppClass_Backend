package routers

import (
	"backend/internal/containers"
	"backend/pkg/middleware"

	"github.com/gin-gonic/gin"
)



func NewAppRouter(app *containers.AppContainer) *gin.Engine {
	router := gin.Default()

	user := router.Group("/user")
	{
		user.POST("/signup",app.UserHandler.SignUp)
		user.POST("/login",app.UserHandler.Login)
	}

	classes := router.Group("/classes")
	classes.Use(middleware.AuthMiddleware())
	{
		classes.POST("/classes",app.ClassesHandler.ViewUserClasses)
		classes.POST("/schedule",app.ClassesHandler.ViewUserSchedule)
	}

	return router
}