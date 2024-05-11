package config

import (
	"github.com/KAI-Back-end/Blog/internal/api/server"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

const (
	envPath = ".env"
)

type Config struct {
	App    App
	Server server.Config
}

type App struct {
	Name    string `yaml:"name"`
	Version string `yaml:"ver"`
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load(envPath); err != nil {
		return nil, err
	}

	cfgApp := new(Config)

	if err := cleanenv.ReadConfig(os.Getenv("GONFIG_PATH"), cfgApp); err != nil {
		return nil, err
	}

	return cfgApp, nil
}
