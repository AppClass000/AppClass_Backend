package routers

import (
	"github.com/gin-gonic/gin"
	"backend/internal/handlers"
	"backend/pkg/middleware"
)

func NewUserRouter(handler handlers.UserHandler) *gin.Engine {

	router := gin.Default()
    router.POST("/signup",handler.SignUp)
	router.POST("/login",handler.Login)
	
	
	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware())

	
	return router
}
