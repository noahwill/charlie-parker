package main

import (
	"charlie-parker/internal/config"
	"charlie-parker/internal/server"
)

func main() {
	config.ConnectRatesTable()
	config.ConnectRouteMetricsTable()
	server.Start()
}
