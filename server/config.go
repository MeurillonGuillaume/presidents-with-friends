package main

import (
	"github.com/koding/multiconfig"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Port int `required:"true"`
}

func MustLoadConfig() Config {
	var cfg Config
	m := multiconfig.New()

	if err := m.Load(&cfg); err != nil {
		logrus.Fatalf("Could not load configuration: %v", err)
	}
	m.MustLoad(&cfg)
	return cfg
}
