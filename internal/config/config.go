package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
)

type Config struct {
	Env            string     `json:"env" env-default:"local"`
	GRPC           GRPCConfig `json:"grpc"`
	MigrationsPath string
	Redis          Redis    	`json:"redis"`
	Database       Database 	`json:"database"`
}

type GRPCConfig struct {
	Port    int           `json:"port"`
	Timeout time.Duration `json:"timeout"`
}

func NewConfig() (*Config, error) {
	var config Config
	if err := config.readConfig(); err != nil {
		return nil, err
	}

	return &config, nil
}

func (config *Config) readConfig() error {
	configPath := os.Getenv("CONFIG_PATH")
	if strings.TrimSpace(configPath) == "" {
		configPath = "./config/local.json"
	}


	fmt.Printf("Using CONFIG_PATH: %s\n", configPath)

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return errors.Errorf("config file does not exist - `%s`", configPath)
	}

	err := cleanenv.ReadConfig(configPath, config)
	if err != nil {
		return errors.Wrapf(err, "failed to load config from `%s`", configPath)
	}

	fmt.Printf("Configuration loaded successfully from: %s\n", configPath)
	return nil
}

