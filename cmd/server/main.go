package main

import (

	"backend/internal/containers"
	"backend/internal/database"
	"backend/internal/routers"
	"log"

)

func main() {
	database.InitDB()
	app := containers.NewAppContainer()

	router := routers.NewAppRouter(app)

	port := ":8080"
	if err := router.Run(port); err != nil {
		log.Fatalf("サーバー起動に失敗: %v", err)
	}

}
