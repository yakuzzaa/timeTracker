package config

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

const (
	envLocal = "local"
	envDev   = "dev"
)

func SetupLogger(env, logPath string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		fileDev, err := os.Create(filepath.Join(logPath, "dev.log"))
		if err != nil {
			fmt.Println(err)
		}
		log = slog.New(
			slog.NewJSONHandler(fileDev, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	}
	return log
}
