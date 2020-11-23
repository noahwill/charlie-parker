package types

// Rate represents a parking rate for a specific Day/Time range
type Rate struct {
	UUID  string `dynamo:"UUID,hash" json:"UUID"`
	Days  string `dynamo:"Days" json:"days"`
	Times string `dynamo:"Times" json:"times"`
	TZ    string `dynamo:"TZ" json:"tz"`
	Price int    `dynamo:"Price" json:"price"`
}

// GetRatesOutput is the output from the GetAllRatesRoute
type GetRatesOutput struct {
	BaseOutput
	Rates []Rate `json:"rates"`
}

// CreateRateInput is the input to the CreateRateRoute and contains
// the fields necessary to create a new rate
type CreateRateInput struct {
	Days  string `json:"days"`
	Times string `json:"times"`
	TZ    string `json:"tz"`
	Price int    `json:"price"`
}

// CreateRateOutput is the output from the CreateRateRoute
type CreateRateOutput struct {
	BaseOutput
	Rate Rate `json:"rate"`
}

// OverwriteRatesInput is the input to the OverwriteRatesRoute
type OverwriteRatesInput struct {
	Rates *[]CreateRateInput `json:"rates"`
}

// OverwriteRatesOutput is the output from the OverwriteRatesRoute
type OverwriteRatesOutput struct {
	BaseOutput
	Rates []Rate `json:"rates"`
}

// GetTimespanPriceInput is the input to the CalculateTimeSpanCostRoute
type GetTimespanPriceInput struct {
	Start *string `json:"start"`
	End   *string `json:"end"`
}

// GetTimespanPriceOutput is the output from the CalculateTimeSpanCostRoute
type GetTimespanPriceOutput struct {
	BaseOutput
	Price string `json:"price"`
}
