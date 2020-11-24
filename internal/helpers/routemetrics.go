package helpers

import (
	"charlie-parker/internal/config"
	"charlie-parker/pkg/types"
	"time"

	"github.com/labstack/gommon/log"
)

// GetAllRouteMetrics gets all the route metrics from the DB
func GetAllRouteMetrics() ([]types.RouteMetrics, error) {
	var metrics []types.RouteMetrics
	err := config.Config.RouteMetricsTableConn.Scan().All(&metrics)
	return metrics, err
}

// UpdateRouteResponseTime updates the route metrics with the given name
func UpdateRouteResponseTime(start time.Time, routeName string) {
	elapsed := time.Since(start)
	metrics, exists, err := getRouteMetrics(routeName)
	if err != nil {
		log.Errorf("Could not update %s average response time with error: %v", routeName, err)
		return
	}

	if !exists {
		metrics = createRouteMetrics(routeName)
	} else {
		metrics.LastUpdated = time.Now().Unix()
	}

	metrics.HitCount++
	metrics.AvgResponseTime = calculateAvgResponseTime(elapsed, metrics.AvgResponseTime, int64(metrics.HitCount))

	if err = config.Config.RouteMetricsTableConn.Put(&metrics).Run(); err != nil {
		log.Errorf("Could not update %s average response time with error: %v", routeName, err)
		return
	}

	log.Infof("Successfully updated %s average response time!", routeName)
	return
}

// UpdateRouteSuccessFailureCount updates a route's success failure count
func UpdateRouteSuccessFailureCount(success bool, routeName string) {
	metrics, exists, err := getRouteMetrics(routeName)
	if err != nil {
		log.Errorf("Could not update %s success/failure count with error: %v", routeName, err)
		return
	}

	if !exists {
		metrics = createRouteMetrics(routeName)
	} else {
		metrics.LastUpdated = time.Now().Unix()
	}

	if success {
		metrics.SuccessCount++
	} else {
		metrics.FailureCount++
	}
	metrics.LastUpdated = time.Now().Unix()

	if err = config.Config.RouteMetricsTableConn.Put(&metrics).Run(); err != nil {
		log.Errorf("Could not update %s success/failure count with error: %v", routeName, err)
		return
	}

	log.Infof("Successfully updated %s success/failure count!", routeName)
	return
}
