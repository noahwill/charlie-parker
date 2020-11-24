package types

// RouteMetrics represents pertinent data around the health of this service's api routes
type RouteMetrics struct {
	RouteName       string `dynamo:"RouteName,omitempty" json:"routeName,omitempty"`
	UUID            string `dynamo:"UUID,hash,omitempty" json:"UUID,omitempty"`
	LastUpdated     int64  `dynamo:"LastUpdated,omitempty" json:"lastUpdated,omitempty"`
	CreatedAt       int64  `dynamo:"CreatedAt,omitempty" json:"createdAt,omitempty"`
	AvgResponseTime string `dynamo:"AvgResponseTime,omitempty" json:"avgResponseTime,omitempty"`
	HitCount        int    `dynamo:"HitCount,omitempty" json:"hitCount,omitempty"`
	SuccessCount    int    `dynamo:"SuccessCount,omitempty" json:"successCount,omitempty"`
	FailureCount    int    `dynamo:"FailureCount,omitempty" json:"failureCount,omitempty"`
}

// GetAllRouteMetricsOutput is the ouput from the GetAllRouteMetricsOutput
type GetAllRouteMetricsOutput struct {
	BaseOutput
	AllRouteMetrics []RouteMetrics `json:"allRouteMetrics"`
}
