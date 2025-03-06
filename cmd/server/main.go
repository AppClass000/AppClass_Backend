package main

import (

	"backend/internal/containers"
	"backend/internal/database"
	"backend/internal/routers"
	"log"

	"github.com/gin-contrib/cors"
)

func main() {
	database.InitDB()
    
	app := containers.NewAppContainer()

	router := routers.NewAppRouter(app)

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

	
	port := ":8080"
	if err := router.Run(port); err != nil {
		log.Fatalf("サーバー起動に失敗: %v", err)
	}

}
