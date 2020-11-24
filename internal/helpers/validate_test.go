package helpers

import (
	"charlie-parker/pkg/types"
	"testing"
	"time"
)

func Test_validateCreateRateInput(t *testing.T) {
	tests := []struct {
		name    string
		in      *types.CreateRateInput
		wantErr bool
	}{
		{
			name: "Simple Passing Validation",
			in: &types.CreateRateInput{
				Days:  "mon,tues,thurs",
				Times: "0900-2100",
				TZ:    "America/Chicago",
				Price: 1500,
			},
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := validateCreateRateInput(test.in, false); (err != nil) != test.wantErr {
				t.Errorf("validateCreateRateInput() error = %v, wantErr %v", err, test.wantErr)
				return
			}
		})
	}
}

func Test_validateAgainstExistingRates(t *testing.T) {
	tests := []struct {
		name          string
		existingRates []types.Rate
		in            types.CreateRateInput
		wantErr       bool
	}{
		{
			name:          "Simple Passing Validation I",
			existingRates: []types.Rate{},
			in: types.CreateRateInput{
				Days:  "mon",
				Times: "0900-1200",
				TZ:    "America/Chicago",
				Price: 1500,
			},
			wantErr: false,
		},
		{
			name: "Simple Passing Validation II",
			existingRates: []types.Rate{
				{
					UUID:  "0000001",
					Days:  "mon",
					Times: "1300-1400",
					TZ:    "America/Chicago",
					Price: 1600,
				},
			},
			in: types.CreateRateInput{
				Days:  "mon",
				Times: "0900-1200",
				TZ:    "America/Chicago",
				Price: 1500,
			},
			wantErr: false,
		},
		{
			name: "Complex Passing Validation",
			existingRates: []types.Rate{
				{
					UUID:  "0000001",
					Days:  "mon,tues,thurs",
					Times: "1300-1400",
					TZ:    "America/Chicago",
					Price: 1600,
				},
			},
			in: types.CreateRateInput{
				Days:  "mon,tues,thurs",
				Times: "0900-1200",
				TZ:    "America/Chicago",
				Price: 1500,
			},
			wantErr: false,
		},
		{
			name: "Simple Overlap Error",
			existingRates: []types.Rate{
				{
					UUID:  "0000001",
					Days:  "mon",
					Times: "0900-1200",
					TZ:    "America/Chicago",
					Price: 1600,
				},
			},
			in: types.CreateRateInput{
				Days:  "mon",
				Times: "0900-1200",
				TZ:    "America/Chicago",
				Price: 1500,
			},
			wantErr: true,
		},
		{
			name: "Complex Overlap Error",
			existingRates: []types.Rate{
				{
					UUID:  "0000001",
					Days:  "mon,tues,thurs",
					Times: "1000-1400",
					TZ:    "America/Chicago",
					Price: 1600,
				},
			},
			in: types.CreateRateInput{
				Days:  "mon,tues,thurs",
				Times: "0900-1200",
				TZ:    "America/Chicago",
				Price: 1500,
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := validateAgainstExistingRates(test.existingRates, test.in); (err != nil) != test.wantErr {
				t.Errorf("validateAgainstExistingRates() error = %v, wantErr %v", err, test.wantErr)
				return
			}
		})
	}
}

func Test_validatePrice(t *testing.T) {
	tests := []struct {
		name    string
		price   int
		wantErr bool
	}{
		{
			name:    "Simple Passing Validation",
			price:   1300,
			wantErr: false,
		},
		{
			name:    "Zero Price",
			price:   0,
			wantErr: true,
		},
		{
			name:    "Negative Price",
			price:   -1,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var err error
			if err = validatePrice(test.price); (err != nil) != test.wantErr {
				t.Errorf("validatePrice() error = %v, wantErr %v", err, test.wantErr)
				return
			}
		})
	}
}

func Test_validateDays(t *testing.T) {
	tests := []struct {
		name    string
		days    string
		wantErr bool
	}{
		{
			name:    "Simple Passing Validation",
			days:    "mon,tues,thurs",
			wantErr: false,
		},
		{
			name:    "Empty Days Error",
			days:    "",
			wantErr: true,
		},
		{
			name:    "Repeating Days Error",
			days:    "mon,mon",
			wantErr: true,
		},
		{
			name:    "Invalid Day Error",
			days:    "INVALID",
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var err error
			if err = validateDays(test.days); (err != nil) != test.wantErr {
				t.Errorf("validateDays() error = %v, wantErr %v", err, test.wantErr)
				return
			}
		})
	}
}

func Test_validateTimeZone(t *testing.T) {
	tests := []struct {
		name    string
		tz      string
		wantErr bool
	}{
		{
			name:    "Simple Passing Validation",
			tz:      "America/Chicago",
			wantErr: false,
		},
		{
			name:    "Empty Time Zone Error",
			tz:      "",
			wantErr: true,
		},
		{
			name:    "Invalid Time Zone Error",
			tz:      "INVALID",
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var err error
			if err = validateTimeZone(test.tz); (err != nil) != test.wantErr {
				t.Errorf("validateTimeZone() error = %v, wantErr %v", err, test.wantErr)
				return
			}
		})
	}
}

func Test_validateTimespan(t *testing.T) {
	tests := []struct {
		name     string
		timespan string
		wantErr  bool
	}{
		{
			name:     "Simple Passing Validation",
			timespan: "0900-1200",
			wantErr:  false,
		},
		{
			name:     "Empty Timespan Error",
			timespan: "",
			wantErr:  true,
		},
		{
			name:     "Bad Format Timespan Error 1",
			timespan: "0900-1200-1500",
			wantErr:  true,
		},
		{
			name:     "Bad Format Timespan Error 2",
			timespan: "09:00-12:00",
			wantErr:  true,
		},
		{
			name:     "Start After End Error",
			timespan: "1200-0900",
			wantErr:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var err error
			if err = validateTimespan(test.timespan); (err != nil) != test.wantErr {
				t.Errorf("validateTimespan() error = %v, wantErr %v", err, test.wantErr)
				return
			}
		})
	}
}

func Test_validateOverlappingRanges(t *testing.T) {
	am0900, _ := time.Parse("1504", "0900")
	am1000, _ := time.Parse("1504", "1000")
	pm1200, _ := time.Parse("1504", "1200")
	pm1300, _ := time.Parse("1504", "1300")
	pm1400, _ := time.Parse("1504", "1400")

	tests := []struct {
		name           string
		existingRanges []timeRange
		newRanges      []timeRange
		wantErr        bool
	}{
		{
			name:           "Simple Passing Validation",
			existingRanges: []timeRange{},
			newRanges: []timeRange{
				{
					earlier: am0900,
					later:   pm1200,
				},
			},
			wantErr: false,
		},
		{
			name: "Passing Validation With Check",
			existingRanges: []timeRange{
				{
					earlier: pm1300,
					later:   pm1400,
				},
			},
			newRanges: []timeRange{
				{
					earlier: am0900,
					later:   pm1200,
				},
			},
			wantErr: false,
		},
		{
			name: "Equal Overlap Error",
			existingRanges: []timeRange{
				{
					earlier: am0900,
					later:   pm1200,
				},
			},
			newRanges: []timeRange{
				{
					earlier: am0900,
					later:   pm1200,
				},
			},
			wantErr: true,
		},
		{
			name: "Soft Overlap Error",
			existingRanges: []timeRange{
				{
					earlier: am1000,
					later:   pm1300,
				},
			},
			newRanges: []timeRange{
				{
					earlier: am0900,
					later:   pm1200,
				},
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var err error
			if err = validateOverlappingRanges(test.existingRanges, test.newRanges); (err != nil) != test.wantErr {
				t.Errorf("validateOverlappingRanges() error = %v, wantErr %v", err, test.wantErr)
				return
			}
		})
	}
}

func Test_validateTimeRange(t *testing.T) {
	tests := []struct {
		name    string
		start   string
		end     string
		wantErr bool
	}{
		{
			name:    "Simple Passing Validation",
			start:   "2015-07-01T07:00:00-05:00",
			end:     "2015-07-01T12:00:00-05:00",
			wantErr: false,
		},
		{
			name:    "Start Parse Error",
			start:   "",
			end:     "2015-07-01T12:00:00-05:00",
			wantErr: true,
		},
		{
			name:    "End Parse Error",
			start:   "2015-07-01T07:00:00-05:00",
			end:     "",
			wantErr: true,
		},
		{
			name:    "Different Year Error",
			start:   "2015-07-01T07:00:00-05:00",
			end:     "2016-07-01T12:00:00-05:00",
			wantErr: true,
		},
		{
			name:    "Different Day Error",
			start:   "2015-07-01T07:00:00-05:00",
			end:     "2015-08-01T12:00:00-05:00",
			wantErr: true,
		},
		{
			name:    "Equal Start/End Error",
			start:   "2015-07-01T07:00:00-05:00",
			end:     "2015-07-01T07:00:00-05:00",
			wantErr: true,
		},
		{
			name:    "Start After End Error",
			start:   "2015-07-01T07:00:00-05:00",
			end:     "2015-07-01T05:00:00-05:00",
			wantErr: true,
		},
		{
			name:    "Different Offset Error",
			start:   "2015-07-01T07:00:00-05:00",
			end:     "2015-07-01T12:00:00-06:00",
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var err error
			if _, _, err = validateTimeRange(&test.start, &test.end); (err != nil) != test.wantErr {
				t.Errorf("validateTimeRange() error = %v, wantErr %v", err, test.wantErr)
				return
			}
		})
	}
}
