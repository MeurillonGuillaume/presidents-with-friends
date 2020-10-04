package main

import (
	"github.com/presidents-with-friends/presidents-with-friends-server/wsserver"
	"github.com/sirupsen/logrus"
)

var (
	version = "0.0.1"
)

func main() {
	logrus.Infof("Presidents-with-friends server v%s", version)

	cfg := MustLoadConfig()

	gameServer := wsserver.NewPresidentServer(cfg.Server)
	gameServer.Start()
}
