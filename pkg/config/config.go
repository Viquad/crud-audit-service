package config

import (
	"github.com/Viquad/crud-audit-service/pkg/database"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

type Config struct {
	DB database.ConnectionInfo `mapstructure:"mongo"`
}

func New(path, name string) (*Config, error) {
	var cfg Config

	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigName(name)

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := envconfig.Process("mongo", &cfg.DB); err != nil {
		return nil, err
	}

	return &cfg, nil
}
