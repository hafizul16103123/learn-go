package config

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env string `env:"ENV" env-default:"dev" validate:"required,oneof=dev test prod"`
}

func MustLoad() (*Config, error) {

	env := os.Getenv("APP_ENV") // $env:APP_ENV='prod'; go run main.go
	if env == "" {
		env = "dev" // fall back to dev when APP_ENV isn't set
	}

	cfg := &Config{}
	file := fmt.Sprintf("%s.env", env) // selects dev.env / prod.env / test.env based on APP_ENV

	if err := cleanenv.ReadConfig(file, cfg); err != nil {
		return nil, err
	}

	//validate
	validator := validator.New()
	if err := validator.Struct(cfg); err != nil {
		return nil, err
	}

	return cfg, nil

}
