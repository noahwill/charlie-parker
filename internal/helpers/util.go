package helpers

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	sun   = "sun"
	mon   = "mon"
	tues  = "tues"
	wed   = "wed"
	thurs = "thurs"
	fri   = "fri"
	sat   = "sat"
)

// isValidDay returns an error for undefined day types
func isValidDay(day string) error {
	switch day {
	case sun, mon, tues, wed, thurs, fri, sat:
		return nil
	}
	return fmt.Errorf("Invalid day: %s", day)
}

// dayToWeekday converts a Day to its corresponding Weekday as defined by time.Weekday
func dayToWeekday(day string) (int, error) {
	switch day {
	case sun:
		return 0, nil
	case mon:
		return 1, nil
	case tues:
		return 2, nil
	case wed:
		return 3, nil
	case thurs:
		return 4, nil
	case fri:
		return 5, nil
	case sat:
		return 6, nil
	}
	return -1, fmt.Errorf("cannot convert %s to type time.Weekday", day)
}

// weekdayToDay converts a Weekday as defined by time.Weekday to its corresponding Day
func weekdayToDay(weekday time.Weekday) (string, error) {
	switch weekday {
	case time.Sunday:
		return sun, nil
	case time.Monday:
		return mon, nil
	case time.Tuesday:
		return tues, nil
	case time.Wednesday:
		return wed, nil
	case time.Thursday:
		return thurs, nil
	case time.Friday:
		return fri, nil
	case time.Saturday:
		return sat, nil
	}
	return "n/a", fmt.Errorf("cannot convert %v to type Day", weekday)
}

// timeRanges are constructed using days/times/timezones from CreateRateInput objects
type timeRange struct {
	days    string
	times   string
	tz      string
	price   int
	earlier time.Time
	later   time.Time
}

// getTimeRangesFromDaysAndTimes generates a list of timeRange objects
// based on the given days and times and tz
func getTimeRangesFromDaysAndTimes(days, times, tz string, price int) []timeRange {
	daysSlice := strings.Split(days, ",")
	timeSpanSlice, _ := timeSpanAsSlice(times)
	loc, _ := time.LoadLocation(tz)

	var timeRanges []timeRange
	for _, day := range daysSlice {
		var timeRange timeRange
		earlier, later, _ := getTimeObjectsFromTimes(timeSpanSlice)
		weekday, _ := dayToWeekday(day)

		// in order to correctly parse the times into a timezone,
		// the year needs to be non-zero - here the default is
		// 2017 since that was a year that started on a sunday
		// and is nice for adding days to get the correct day
		// portion of the time object
		earlier = earlier.AddDate(2017, 0, weekday).In(loc)
		later = later.AddDate(2017, 0, weekday).In(loc)

		timeRange.days = days
		timeRange.times = times
		timeRange.tz = tz
		timeRange.price = price
		timeRange.earlier = earlier
		timeRange.later = later
		timeRanges = append(timeRanges, timeRange)
	}

	return timeRanges
}

// timeSpanAsSlice returns a slice containing two strings representing hours of the day
func timeSpanAsSlice(timespan string) ([]string, error) {
	times := strings.Split(timespan, "-")
	if len(times) != 2 {
		return times, errors.New("times should range between only two hours of the day")
	}
	return times, nil
}

// getTimeObjectsFromTimes returns the earlier and later Time representation of a slice "times" containing two hours
// the times returned have the format "0000-01-01 HH:00:00 +0000 UTC"
func getTimeObjectsFromTimes(times []string) (earlier time.Time, later time.Time, err error) {
	if earlier, err = time.Parse("1504", times[0]); err != nil {
		return earlier, later, fmt.Errorf("could not parse earlier time in range %v: %v", times, err)
	} else if later, err = time.Parse("1504", times[1]); err != nil {
		return earlier, later, fmt.Errorf("could not parse later time in range %v: %v", times, err)
	}
	return earlier, later, err
}
