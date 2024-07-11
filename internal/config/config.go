package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env-default:"local" env-required:"true"`
	HTTPServer `yaml:"http_server"`
	Database   `yaml:"db"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-required:"true"`
	Timeout     time.Duration `yaml:"timeout" env-required:"true"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-required:"true"`
}

type Database struct {
	DBHost     string `yaml:"db_host" env-required:"true"`
	DBPort     string `yaml:"db_port" env-required:"true"`
	DBUser     string `yaml:"db_user" env-required:"true"`
	DBPassword string `yaml:"db_password" env-required:"true"`
	DBName     string `yaml:"db_name" env-required:"true"`
}

func (c *Database) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
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
