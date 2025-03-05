package main

import (
	"backend/internal/app/repositories"
	"backend/internal/app/servises"
	"backend/internal/database"
	"backend/internal/handlers"
	"backend/internal/routers"
	"log"
	"github.com/gin-contrib/cors"
)

func main() {
	database.InitDB()
    userReposirory := repositories.NewUserRepository(database.DB)
	userServise :=servises.NewUserServise(&userReposirory)
	userHandler := handlers.NewUserHandler(userServise)

	router:= routers.NewUserRouter(userHandler)


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
