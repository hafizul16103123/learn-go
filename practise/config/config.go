package config

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env string `env:"ENV" validate:"required,oneof=dev test prod"`
}

func MustLoad() (*Config, error) {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "dev"
	}

	cfg := &Config{}
	file := fmt.Sprintf("%s.env", appEnv)

	if err := cleanenv.ReadConfig(file, cfg); err != nil {
		return nil, err
	}

	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
