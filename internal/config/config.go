package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env:"ENV" default:"local" required:"true"`
	LogPath    string `yaml:"log_path" env:"LOG_PATH" default:"logs/" required:"true"`
	HTTPServer `yaml:"http_server"`
	Database   `yaml:"db"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env:"HTTP_SERVER_ADDRESS" required:"true"`
	Timeout     time.Duration `yaml:"timeout" env:"HTTP_SERVER_TIMEOUT" required:"true"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"HTTP_SERVER_IDLE_TIMEOUT" required:"true"`
}

type Database struct {
	DBHost     string `yaml:"db_host" env:"DB_HOST" required:"true"`
	DBPort     string `yaml:"db_port" env:"DB_PORT" required:"true"`
	DBUser     string `yaml:"db_user" env:"DB_USER" required:"true"`
	DBPassword string `yaml:"db_password" env:"DB_PASSWORD" required:"true"`
	DBName     string `yaml:"db_name" env:"DB_NAME" required:"true"`
	Driver     string `yaml:"driver" env:"DB_DRIVER" required:"true"`
	Dir        string `yaml:"dir" env:"DB_MIGRATIONS_DIR" required:"true"`
}

func (db *Database) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		db.DBHost, db.DBPort, db.DBUser, db.DBPassword, db.DBName)
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("Config path not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file %s does not exist", configPath)
	}

	var config Config

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("Error reading config: %s", err)
	}

	return &config
}
