package main

import (
	"log"
	"timeTracker/internal/config"
	"timeTracker/internal/storage"

	"github.com/pressly/goose/v3"
)

func main() {
	configLoad := config.MustLoad()

	db, err := storage.Connect(configLoad)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err.Error())
	}
	log.Println("Database connected successfully")

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB: %v", err)
	}

	if err := goose.Up(sqlDB, configLoad.Dir); err != nil {
		log.Fatalf("failed to apply migrations: %v", err)
	}

	log.Println("Migrations applied successfully")

	//handlers := handler.NewHandler(authClient, listClient, itemClient)
	//srv := new(api.Server)
	//if err := srv.Run(configLoad.Address, handlers.InitRoutes()); err != nil {
	//	log.Fatalf("Something went wrong: %s", err)
	//}
}
