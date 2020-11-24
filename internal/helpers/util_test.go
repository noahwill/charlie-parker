package helpers

import (
	"charlie-parker/pkg/types"
	"reflect"
	"testing"
	"time"
)

func Test_isValidDay(t *testing.T) {
	tests := []struct {
		name    string
		day     string
		wantErr bool
	}{
		{
			name:    "Sunday Validation",
			day:     sun,
			wantErr: false,
		},
		{
			name:    "Monday Validation",
			day:     mon,
			wantErr: false,
		},
		{
			name:    "Tuesday Validation",
			day:     tues,
			wantErr: false,
		},
		{
			name:    "Wednesday Validation",
			day:     wed,
			wantErr: false,
		},
		{
			name:    "Thursday Validation",
			day:     thurs,
			wantErr: false,
		},
		{
			name:    "Friday Validation",
			day:     fri,
			wantErr: false,
		},
		{
			name:    "Saturday Validation",
			day:     sat,
			wantErr: false,
		},
		{
			name:    "Undefined Day Error",
			day:     "UNDEFINED",
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := isValidDay(test.day); (err != nil) != test.wantErr {
				t.Errorf("isValidDay() error = %v, wantErr %v", err, test.wantErr)
				return
			}
		})
	}
}

func Test_dayToWeekday(t *testing.T) {
	tests := []struct {
		name    string
		day     string
		want    interface{}
		wantErr bool
	}{
		{
			name:    "Sunday Validation",
			day:     sun,
			want:    0,
			wantErr: false,
		},
		{
			name:    "Monday Validation",
			day:     mon,
			want:    1,
			wantErr: false,
		},
		{
			name:    "Tuesday Validation",
			day:     tues,
			want:    2,
			wantErr: false,
		},
		{
			name:    "Wednesday Validation",
			day:     wed,
			want:    3,
			wantErr: false,
		},
		{
			name:    "Thursday Validation",
			day:     thurs,
			want:    4,
			wantErr: false,
		},
		{
			name:    "Friday Validation",
			day:     fri,
			want:    5,
			wantErr: false,
		},
		{
			name:    "Saturday Validation",
			day:     sat,
			want:    6,
			wantErr: false,
		},
		{
			name:    "Undefined Day Error",
			day:     "UNDEFINED",
			want:    -1,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := dayToWeekday(test.day)
			if (err != nil) != test.wantErr {
				t.Errorf("dayToWeekday() error = %v, wantErr %v", err, test.wantErr)
				return
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("dayToWeekday() got = %v, want %v", got, test.want)
			}
		})
	}
}

func Test_weekdayToDay(t *testing.T) {
	tests := []struct {
		name    string
		weekday time.Weekday
		want    interface{}
		wantErr bool
	}{
		{
			name:    "Sunday Validation",
			weekday: time.Sunday,
			want:    sun,
			wantErr: false,
		},
		{
			name:    "Monday Validation",
			weekday: time.Monday,
			want:    mon,
			wantErr: false,
		},
		{
			name:    "Tuesday Validation",
			weekday: time.Tuesday,
			want:    tues,
			wantErr: false,
		},
		{
			name:    "Wednesday Validation",
			weekday: time.Wednesday,
			want:    wed,
			wantErr: false,
		},
		{
			name:    "Thursday Validation",
			weekday: time.Thursday,
			want:    thurs,
			wantErr: false,
		},
		{
			name:    "Friday Validation",
			weekday: time.Friday,
			want:    fri,
			wantErr: false,
		},
		{
			name:    "Saturday Validation",
			weekday: time.Saturday,
			want:    sat,
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := weekdayToDay(test.weekday)
			if (err != nil) != test.wantErr {
				t.Errorf("weekdayToDay() error = %v, wantErr %v", err, test.wantErr)
				return
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("weekdayToDay() got = %v, want %v", got, test.want)
			}
		})
	}
}

func Test_getTimeRangesFromDaysAndTimes(t *testing.T) {
	chi, _ := time.LoadLocation("America/Chicago")
	tests := []struct {
		name  string
		days  string
		times string
		tz    string
		price int
		want  []timeRange
	}{
		{
			name:  "Simple Passing Validation",
			days:  "mon",
			times: "0900-1200",
			tz:    "America/Chicago",
			price: 1500,
			want: []timeRange{
				{
					days:    "mon",
					times:   "0900-1200",
					tz:      "America/Chicago",
					price:   1500,
					earlier: time.Date(2017, time.January, 2, 9, 0, 0, 0, chi),
					later:   time.Date(2017, time.January, 2, 12, 0, 0, 0, chi),
				},
			},
		},
		{
			name:  "Complex Passing Validation",
			days:  "mon,tues,sat",
			times: "0900-1200",
			tz:    "America/Chicago",
			price: 1500,
			want: []timeRange{
				{
					days:    "mon,tues,sat",
					times:   "0900-1200",
					tz:      "America/Chicago",
					price:   1500,
					earlier: time.Date(2017, time.January, 2, 9, 0, 0, 0, chi),
					later:   time.Date(2017, time.January, 2, 12, 0, 0, 0, chi),
				},
				{
					days:    "mon,tues,sat",
					times:   "0900-1200",
					tz:      "America/Chicago",
					price:   1500,
					earlier: time.Date(2017, time.January, 3, 9, 0, 0, 0, chi),
					later:   time.Date(2017, time.January, 3, 12, 0, 0, 0, chi),
				},
				{
					days:    "mon,tues,sat",
					times:   "0900-1200",
					tz:      "America/Chicago",
					price:   1500,
					earlier: time.Date(2017, time.January, 7, 9, 0, 0, 0, chi),
					later:   time.Date(2017, time.January, 7, 12, 0, 0, 0, chi),
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := getTimeRangesFromDaysAndTimes(test.days, test.times, test.tz, test.price)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("getTimeRangesFromDaysAndTimes() got = %v, want %v", got, test.want)
				return
			}
		})
	}
}

func Test_timeSpanAsSlice(t *testing.T) {
	tests := []struct {
		name     string
		timespan string
		want     interface{}
		wantErr  bool
	}{
		{
			name:     "Simple Passing Validation",
			timespan: "0900-1200",
			want:     []string{"0900", "1200"},
			wantErr:  false,
		},
		{
			name:     "Bad Format Timespan Error",
			timespan: "0900-1200-1500",
			want:     []string{"0900", "1200", "1500"},
			wantErr:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := timeSpanAsSlice(test.timespan)
			if (err != nil) != test.wantErr {
				t.Errorf("timeSpanAsSlice() error = %v, wantErr %v", err, test.wantErr)
				return
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("timeSpanAsSlice() got = %v, want %v", got, test.want)
				return
			}
		})
	}
}

func Test_getTimeObjectsFromTimes(t *testing.T) {
	tests := []struct {
		name        string
		times       []string
		earlierWant time.Time
		laterWant   time.Time
		wantErr     bool
	}{
		{
			name:        "Simple Passing Validation",
			times:       []string{"0900", "1200"},
			earlierWant: time.Date(0, time.January, 1, 9, 0, 0, 0, time.UTC),
			laterWant:   time.Date(0, time.January, 1, 12, 0, 0, 0, time.UTC),
			wantErr:     false,
		},
		{
			name:    "Earlier Parsing Error",
			times:   []string{"", "1200"},
			wantErr: true,
		},
		{
			name:    "Later Parsing Error",
			times:   []string{"0900", ""},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			earlier, later, err := getTimeObjectsFromTimes(test.times)
			if (err != nil) != test.wantErr {
				t.Errorf("getTimeObjectsFromTimes() error = %v, wantErr %v", err, test.wantErr)
				return
			}

			if !test.wantErr {
				if !reflect.DeepEqual(earlier, test.earlierWant) {
					t.Errorf("getTimeObjectsFromTimes() got = %v, want %v", earlier, test.earlierWant)
					return
				}

				if !reflect.DeepEqual(later, test.laterWant) {
					t.Errorf("getTimeObjectsFromTimes() got = %v, want %v", later, test.laterWant)
					return
				}
			}

		})
	}
}

func Test_matchTimespanToRate(t *testing.T) {
	chi, _ := time.LoadLocation("America/Chicago")
	tests := []struct {
		name          string
		startTime     time.Time
		endTime       time.Time
		existingRates []types.Rate
		want          interface{}
		wantErr       bool
	}{
		{
			name:      "Exact Match",
			startTime: time.Date(2017, time.January, 2, 9, 0, 0, 0, chi),
			endTime:   time.Date(2017, time.January, 2, 12, 0, 0, 0, chi),
			existingRates: []types.Rate{
				{
					UUID:  "0000001",
					Days:  "mon",
					Times: "0900-1200",
					TZ:    "America/Chicago",
					Price: 1600,
				},
			},
			want: types.Rate{
				UUID:  "0000001",
				Days:  "mon",
				Times: "0900-1200",
				TZ:    "America/Chicago",
				Price: 1600,
			},
			wantErr: false,
		},
		{
			name:      "Within Range Match",
			startTime: time.Date(2017, time.January, 2, 10, 0, 0, 0, chi),
			endTime:   time.Date(2017, time.January, 2, 10, 30, 0, 0, chi),
			existingRates: []types.Rate{
				{
					UUID:  "0000001",
					Days:  "mon",
					Times: "0900-1200",
					TZ:    "America/Chicago",
					Price: 1600,
				},
			},
			want: types.Rate{
				UUID:  "0000001",
				Days:  "mon",
				Times: "0900-1200",
				TZ:    "America/Chicago",
				Price: 1600,
			},
			wantErr: false,
		},
		{
			name:      "Complex Match",
			startTime: time.Date(2017, time.January, 2, 10, 0, 0, 0, chi),
			endTime:   time.Date(2017, time.January, 2, 10, 30, 0, 0, chi),
			existingRates: []types.Rate{
				{
					UUID:  "0000001",
					Days:  "mon,tues,wed",
					Times: "0900-1200",
					TZ:    "America/Chicago",
					Price: 1600,
				},
				{
					UUID:  "0000002",
					Days:  "thurs,fri",
					Times: "0900-1200",
					TZ:    "America/Chicago",
					Price: 1600,
				},
			},
			want: types.Rate{
				UUID:  "0000001",
				Days:  "mon,tues,wed",
				Times: "0900-1200",
				TZ:    "America/Chicago",
				Price: 1600,
			},
			wantErr: false,
		},
		{
			name:      "Multiple Match Error",
			startTime: time.Date(2017, time.January, 2, 10, 0, 0, 0, chi),
			endTime:   time.Date(2017, time.January, 2, 10, 30, 0, 0, chi),
			existingRates: []types.Rate{
				{
					UUID:  "0000001",
					Days:  "mon,tues,wed",
					Times: "0900-1200",
					TZ:    "America/Chicago",
					Price: 1600,
				},
				{
					UUID:  "0000002",
					Days:  "mon,tues,wed",
					Times: "1000-1300",
					TZ:    "America/Chicago",
					Price: 1600,
				},
			},
			wantErr: true,
		},
		{
			name:      "No Match Error",
			startTime: time.Date(2017, time.January, 2, 7, 0, 0, 0, chi),
			endTime:   time.Date(2017, time.January, 2, 7, 30, 0, 0, chi),
			existingRates: []types.Rate{
				{
					UUID:  "0000001",
					Days:  "mon,tues,wed",
					Times: "0900-1200",
					TZ:    "America/Chicago",
					Price: 1600,
				},
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := matchTimespanToRate(test.startTime, test.endTime, test.existingRates)
			if (err != nil) != test.wantErr {
				t.Errorf("matchTimespanToRate() error = %v, wantErr %v", err, test.wantErr)
				return
			}

			if !test.wantErr {
				if !reflect.DeepEqual(got, test.want) {
					t.Errorf("matchTimespanToRate() got = %v, want %v", got, test.want)
					return
				}
			}
		})
	}
}

func Test_getLocationOffset(t *testing.T) {
	chi, _ := time.LoadLocation("America/Chicago")
	tests := []struct {
		name  string
		year  int
		month time.Month
		day   int
		loc   *time.Location
		want  interface{}
	}{
		{
			name:  "Chicago CST",
			year:  2017,
			month: time.March,
			day:   11,
			loc:   chi,
			want:  -21600,
		},
		{
			name:  "Chicago CDT",
			year:  2017,
			month: time.March,
			day:   12,
			loc:   chi,
			want:  -18000,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := getLocationOffset(test.year, test.month, test.day, test.loc)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("getLocationOffset() got = %v, want %v", got, test.want)
				return
			}
		})
	}
}

func Test_isValidRouteName(t *testing.T) {
	tests := []struct {
		name      string
		routeName string
		wantErr   bool
	}{
		{
			name:      "GetRatesRoute Validation",
			routeName: GetRatesRouteName,
			wantErr:   false,
		},
		{
			name:      "CreateRateRoute Validation",
			routeName: CreateRateRouteName,
			wantErr:   false,
		},
		{
			name:      "OverwriteRatesRoute Validation",
			routeName: OverwriteRatesRouteName,
			wantErr:   false,
		},
		{
			name:      "GetTimespanPriceRoute Validation",
			routeName: GetTimespanPriceRouteName,
			wantErr:   false,
		},
		{
			name:      "GetAllRouteMetricsRoute Validation",
			routeName: GetAllRouteMetricsRouteName,
			wantErr:   false,
		},
		{
			name:      "Undefined Error",
			routeName: "",
			wantErr:   true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := isValidRouteName(test.routeName); (err != nil) != test.wantErr {
				t.Errorf("isValidRouteName() error = %v, wantErr %v", err, test.wantErr)
				return
			}
		})
	}
}

func Test_calculateAvgResponseTime(t *testing.T) {
	ms10, _ := time.ParseDuration("10ms")
	tests := []struct {
		name            string
		responseTime    time.Duration
		avgResponseTime string
		hits            int64
		want            interface{}
	}{
		{
			name:            "Simple Pass",
			responseTime:    ms10,
			avgResponseTime: "0ms",
			hits:            1,
			want:            "10ms",
		},
		{
			name:            ">1 Hits Pass",
			responseTime:    ms10,
			avgResponseTime: "39ms",
			hits:            13,
			want:            "37ms",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := calculateAvgResponseTime(test.responseTime, test.avgResponseTime, test.hits)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("calculateAvgResponseTime() got = %v, want %v", got, test.want)
				return
			}
		})
	}
}
