package seeder

import (
	"charlie-parker/internal/helpers"
	"charlie-parker/pkg/types"
	"time"
)

var rateSeed = []types.CreateRateInput{
	{
		Days:  "mon,tues,thurs",
		Times: "0900-2100",
		TZ:    "America/Chicago",
		Price: 1500,
	},
	{
		Days:  "fri,sat,sun",
		Times: "0900-2100",
		TZ:    "America/Chicago",
		Price: 2000,
	},
	{
		Days:  "wed",
		Times: "0600-1800",
		TZ:    "America/Chicago",
		Price: 1750,
	},
	{
		Days:  "mon,wed,sat",
		Times: "0100-0500",
		TZ:    "America/Chicago",
		Price: 1000,
	},
	{
		Days:  "sun,tues",
		Times: "0100-0700",
		TZ:    "America/Chicago",
		Price: 925,
	},
}

var routeMetricsSeed = []types.RouteMetrics{
	{
		UUID:            "81203de7-0fa8-40a3-8927-bae779c036e3",
		RouteName:       helpers.GetRatesRouteName,
		CreatedAt:       time.Now().Unix(),
		LastUpdated:     time.Now().Unix(),
		AvgResponseTime: "0s",
	},
	{
		UUID:            "8e4fdd48-2ac0-41d5-9329-d61bfb5b5ffa",
		RouteName:       helpers.CreateRateRouteName,
		CreatedAt:       time.Now().Unix(),
		LastUpdated:     time.Now().Unix(),
		AvgResponseTime: "0s",
	},
	{
		UUID:            "3107144a-08c6-47be-a500-74cc03348f7f",
		RouteName:       helpers.OverwriteRatesRouteName,
		CreatedAt:       time.Now().Unix(),
		LastUpdated:     time.Now().Unix(),
		AvgResponseTime: "0s",
	},
	{
		UUID:            "623bc8e5-330a-428f-b906-41e2d18293ca",
		RouteName:       helpers.GetTimespanPriceRouteName,
		CreatedAt:       time.Now().Unix(),
		LastUpdated:     time.Now().Unix(),
		AvgResponseTime: "0s",
	},
	{
		UUID:            "5699741b-373c-4055-99f0-d5d883785ec2",
		RouteName:       helpers.GetAllRouteMetricsRouteName,
		CreatedAt:       time.Now().Unix(),
		LastUpdated:     time.Now().Unix(),
		AvgResponseTime: "0s",
	},
}
