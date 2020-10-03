package main

import "github.com/sirupsen/logrus"

var (
	version = "0.0.1"
)

func main() {
	logrus.Infof("Presidents-with-friends server v%s", version)

	cfg := MustLoadConfig()
	logrus.Info(cfg)
}
