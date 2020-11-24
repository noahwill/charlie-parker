package main

import (
	"charlie-parker/internal/config"
	"charlie-parker/internal/seeder"

	"github.com/labstack/gommon/log"
)

func main() {
	config.ConnectRatesTable()
	config.ConnectRouteMetricsTable()
	log.Infof("%s starting", config.Config.AppName)
	seeder.Run()
}
