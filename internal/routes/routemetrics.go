package routes

import (
	"charlie-parker/internal/config"
	"charlie-parker/internal/helpers"
	"charlie-parker/pkg/types"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// GetAllRouteMetricsRoute is the api handler for getting all charlie-parker
// route metrics from the DB
func GetAllRouteMetricsRoute(c echo.Context) error {
	defer helpers.UpdateRouteResponseTime(time.Now(), helpers.GetAllRouteMetricsRouteName)
	var (
		err          error
		foundMetrics []types.RouteMetrics
		out          types.GetAllRouteMetricsOutput
	)

	if foundMetrics, err = helpers.GetAllRouteMetrics(); err != nil {
		out.Error = fmt.Sprintf("Could not get all route metrics from %s with error: %v", config.Config.RouteMetricsTable, err)
		log.Error(out.Error)
		defer helpers.UpdateRouteSuccessFailureCount(false, helpers.GetAllRouteMetricsRouteName)
		return c.JSON(http.StatusInternalServerError, &out)
	}

	out.Ok = true
	out.AllRouteMetrics = foundMetrics
	log.Infof("Successfully got all %d rates from %s", len(out.AllRouteMetrics), config.Config.RouteMetricsTable)
	defer helpers.UpdateRouteSuccessFailureCount(true, helpers.GetAllRouteMetricsRouteName)
	return c.JSON(http.StatusOK, &out)
}
