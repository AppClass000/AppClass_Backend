package main

import (
	"backend/database"
	"backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	database.InitDB()
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

	api := router.Group("/api", routes.AuthMiddleware())

	routes.SignUproute(router)
	routes.ResponseJWTroute(router)
	routes.GetUserClasssesroute(api)
	routes.UserDetailroute(api)
	routes.GetScheduleDataroute(api)
	routes.RegisterClassesroute(api)

	port := ":8080"
	if err := router.Run(port); err != nil {
		log.Fatalf("サーバー起動に失敗: %v", err)
	}

}
