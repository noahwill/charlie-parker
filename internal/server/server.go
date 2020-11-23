package server

import (
	"time"

	"charlie-parker/internal/config"
	utils "charlie-parker/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// Start will start the api server
func Start() {
	log.Info("Server up and running")
	e := echo.New()

	// TODO: Add some better defaults for the http server
	// to prevent connection starving
	e.Server.IdleTimeout = 30 * time.Second
	e.Server.ReadTimeout = 15 * time.Second
	e.Server.ReadHeaderTimeout = 10 * time.Second

	// Utility
	e.GET("/heartbeat", utils.HeartbeatRoute)

	e.Logger.Fatal(e.Start(":" + config.Config.WebServerPort))
}
