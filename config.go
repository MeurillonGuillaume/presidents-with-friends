package main

import (
	"github.com/koding/multiconfig"
	"github.com/presidents-with-friends/presidents-with-friends/server/wsserver"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Server wsserver.PresidentConfig
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
