package server

import (
	"time"

	"charlie-parker/internal/config"
	"charlie-parker/internal/routes"
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

	// V1 API route group
	v1 := e.Group("/api/v1")
	// RATES
	v1.GET("/rates", routes.GetRatesRoute)
	v1.POST("/rates/create", routes.CreateRateRoute)
	v1.POST("/rates/update/all", routes.OverwriteRatesRoute)
	// PARKING PRICE
	v1.POST("/park", routes.GetTimespanPriceRoute)

	e.Logger.Fatal(e.Start(":" + config.Config.WebServerPort))
}
