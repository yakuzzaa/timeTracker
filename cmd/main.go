package main

import (
	"log"
	_ "timeTracker/docs"
	"timeTracker/internal/api"
	"timeTracker/internal/api/handlers"
	"timeTracker/internal/api/repository"
	sv "timeTracker/internal/api/services"
	"timeTracker/internal/config"
	"timeTracker/internal/storage"

	"github.com/pressly/goose/v3"
)

// @title TimeTracker API
// @version 1.0
// @description API for time tracking app
// @host localhost:8080
// @BasePath /api
func main() {
	configLoad := config.MustLoad()

	db, err := storage.Connect(configLoad)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err.Error())
	}
	log.Println("Database connected successfully")

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB: %v", err)
	}

	if err := goose.Up(sqlDB, configLoad.Dir); err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	log.Println("Migrations applied successfully")

	repos := repository.NewRepository(db)
	services := sv.NewService(repos)
	handler := handlers.NewHandler(services)

	srv := new(api.Server)
	if err := srv.Run(configLoad.Address, handler.InitRoutes()); err != nil {
		log.Fatalf("Something went wrong: %s", err)
	}
}
