package configs

import (
	"time"

	"github.com/go-playground/validator"
)

type Config struct {
	App      App
	Log      Log
	Database Database
	Redis    Redis
	Secrets  Secrets
}

func (c Config) Validate() error {
	return validator.New().Struct(c)
}

type App struct {
	Name    string        `validate:"required"`
	Port    string        `validate:"required"`
	Timeout time.Duration `validate:"gt=0"`
}

type Log struct {
	Level string `validate:"required"`
	Env   string `validate:"required"`
}

type Database struct {
	Host     string `validate:"required"`
	Port     string `validate:"required"`
	Database string `validate:"required"`
}

type Redis struct {
	Addr     string `validate:"required"`
	Port     string `validate:"required"`
	DB       int
	PoolSize int `validate:"required"`
}

type Secrets struct {
	RedisPassword string
	DbUsername    string
	DbPassword    string
}
