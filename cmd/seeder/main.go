package main

import (
	"charlie-parker/internal/config"
	"charlie-parker/internal/seeder"

	"github.com/labstack/gommon/log"
)

func main() {
	log.Infof("%s starting", config.Config.AppName)
	seeder.Run()
}
