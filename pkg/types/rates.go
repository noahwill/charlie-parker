package types

// Rate represents a parking rate for a specific Day/Time range
type Rate struct {
	UUID  string `dynamo:"UUID,hash" json:"UUID"`
	Days  string `dynamo:"Days" json:"days"`
	Times string `dynamo:"Times" json:"times"`
	TZ    string `dynamo:"TZ" json:"tz"`
	Price int    `dynamo:"Price" json:"price"`
}
