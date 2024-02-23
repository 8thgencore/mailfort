package config

import (
	"fmt"
	"log/slog"
	"path"
	"time"

	"github.com/8thgencore/mailfort/pkg/logger/slogpretty"
	"github.com/ilyakaznacheev/cleanenv"
)

type Env string

const (
	Local Env = "local"
	Dev   Env = "dev"
	Prod  Env = "prod"
)

type (
	Config struct {
		Env  Env  `yaml:"env" env-defaul:"local" env-required:"true"` // local, dev or prod
		App  App  `yaml:"app"`
		HTTP HTTP `yaml:"http"`
		GRPC GRPC `yaml:"grpc"`
		Mail Mail `yaml:"mail"`
		Log  Log  `yaml:"log"`
	}

	// App contains all the environment variables for the application
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP contains all the environment variables for the http server
	HTTP struct {
		Host         string `env-required:"true" yaml:"host" env:"HTTP_HOST"`
		Port         string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
		AllowOrigins string `env-required:"true" yaml:"allow_origins"`
	}

	// GRPC contains all the environment variables for the gRPC server
	GRPC struct {
		Port    int           `env-required:"true" yaml:"port"`
		Timeout time.Duration `env-required:"true" yaml:"timeout"`
	}

	// Mail contains all the environment variables for the mail client
	Mail struct {
		Host     string `env-required:"true" yaml:"host"     env:"MAIL_HOST"`
		Port     int    `env-required:"true" yaml:"port"     env:"MAIL_PORT"`
		Username string `env-required:"true" yaml:"username" env:"MAIL_USERNAME"`
		Password string `env-required:"true" yaml:"password" env:"MAIL_PASSWORD"`
	}

	// Logger settings
	Log struct {
		Slog Slog `yaml:"slog"`
	}
	Slog struct {
		Level     slog.Level              `yaml:"level"`
		AddSource bool                    `yaml:"add_source"`
		Format    slogpretty.FieldsFormat `yaml:"format"` // json, text or pretty
		Pretty    PrettyLog               `yaml:"pretty"`
	}
	PrettyLog struct {
		FieldsFormat slogpretty.FieldsFormat `yaml:"fields_format"` // json, json-indent or yaml
		Emoji        bool                    `yaml:"emoji"`
		TimeLayout   string                  `yaml:"time_layout"`
	}
)

func NewConfig(configPath string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(path.Join("./", configPath), cfg)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	err = cleanenv.UpdateEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("error updating env: %w", err)
	}

	return cfg, nil
}
