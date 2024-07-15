package main

import (
	"log/slog"
	"os"

	_ "github.com/yakuzzaa/timeTracker/docs"
	"github.com/yakuzzaa/timeTracker/internal/api"
	"github.com/yakuzzaa/timeTracker/internal/api/handlers"
	"github.com/yakuzzaa/timeTracker/internal/api/repository"
	sv "github.com/yakuzzaa/timeTracker/internal/api/services"
	"github.com/yakuzzaa/timeTracker/internal/config"

	"github.com/pressly/goose/v3"
)

// @title TimeTracker API
// @version 1.0
// @description API for time tracking app
// @host localhost:8080
// @BasePath /api
func main() {
	configLoad := config.MustLoad()

	logger := config.SetupLogger(configLoad.Env, configLoad.LogPath)
	slog.SetDefault(logger)
	slog.Info("starting server", slog.String("env", configLoad.Env))
	slog.Debug("debug logging enabled")

	db, err := config.DbConnect(configLoad)
	if err != nil {
		slog.Error("Failed to connect to database: %v", err)
		os.Exit(1)
	}
	logger.Info("Database connected successfully")

	sqlDB, err := db.DB()
	if err != nil {
		slog.Error("Failed to get sql.DB: %v", err)
		os.Exit(1)
	}

	if err := goose.Up(sqlDB, configLoad.Dir); err != nil {
		slog.Error("Failed to apply migrations: %v", err)
		os.Exit(1)
	}

	slog.Info("Migrations applied successfully")

	repos := repository.NewRepository(db)

	services := sv.NewService(repos, logger)
	handler := handlers.NewHandler(services, logger)

	srv := new(api.Server)
	if err := srv.Run(configLoad.Address, handler.InitRoutes()); err != nil {
		slog.Info("Something went wrong: %s", err)
	}
}
