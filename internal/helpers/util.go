package helpers

import (
	"charlie-parker/pkg/types"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/labstack/gommon/log"
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

// matchTimespanToRate tries to find an existing rate that a given timespan would be covered by
func matchTimespanToRate(startTime, endTime time.Time) (types.Rate, error) {
	var (
		err           error
		rate          types.Rate
		existingRates []types.Rate
		matchingRates []types.Rate
	)

	if existingRates, err = GetRates(); err != nil {
		return rate, err
	}

	inputDayStr, _ := weekdayToDay(startTime.Weekday())
	inputDay := startTime.Day()
	inputMonth := startTime.Month()
	inputYear := startTime.Year()
	_, inputOffset := startTime.Zone()

	for _, rate := range existingRates {
		if strings.Contains(rate.Days, inputDayStr) {
			rateLocation, _ := time.LoadLocation(rate.TZ)
			// create a dummy time object in terms of the input's year, month, and day
			// set the hour of the time to 3am to widely avoid any potential conflict for days
			// on which clocks change around the world (the latest of which currently occurs
			// the first Sunday in April at 4AM in Samoa)
			dummy := time.Date(inputYear, inputMonth, inputDay, 5, 0, 0, 0, rateLocation)
			_, rateOffset := dummy.Zone()
			if inputOffset == rateOffset {
				rateTimes, _ := timeSpanAsSlice(rate.Times)
				// these times have the format "0000-01-01 HH:00:00 +0000 UTC"
				rateStart, rateEnd, _ := getTimeObjectsFromTimes(rateTimes)
				// put rate start and end in terms of the input's year, month, and day;
				// due to the month and day of the existing rate times being set to 01,
				// the month and day passed in are decremented
				rateStart = rateStart.AddDate(inputYear, int(inputMonth)-1, inputDay-1)
				rateEnd = rateEnd.AddDate(inputYear, int(inputMonth)-1, inputDay-1)
				// rate start and end are still in UTC, thus we must put the times in the correct
				// timezone while retaining the same hour information by subtracting the offset
				// from their unix timestamp representation
				rateStart = time.Unix((rateStart.Unix() - int64(inputOffset)), 0).In(rateLocation)
				rateEnd = time.Unix((rateEnd.Unix() - int64(inputOffset)), 0).In(rateLocation)
				// rateStart <= startTime < rateEnd
				if startTime.Sub(rateStart) >= 0 && startTime.Sub(rateEnd) < 0 {
					// rateStart < endTime <= rateEnd
					if endTime.Sub(rateStart) > 0 && endTime.Sub(rateEnd) <= 0 {
						// We found a match!
						matchingRates = append(matchingRates, rate)
					}
					log.Infof("Input end time (%v) is does not fall between rate start and end (%v - %v)", endTime, rateStart, rateEnd)
				} else {
					log.Infof("Input start time (%v) is does not fall between rate start and end (%v - %v)", startTime, rateStart, rateEnd)
				}
			} else {
				log.Infof("Input offset (%v) not equal to rate offset (%v)", inputOffset, rateOffset)
			}
		} else {
			log.Infof("Input day (%s) not in rate's days %s", inputDay, rate.Days)
		}
	}

	if len(matchingRates) != 1 {
		log.Errorf("There were multiple rates that matched a user's GetTimespanPriceInput, this may indicate that somehow there are rates that overlap in the DB. (Matched Rates: %v)", matchingRates)
		return rate, fmt.Errorf("there was not one rate found to match the given start/end (%v - %v)", startTime, endTime)
	}

	return matchingRates[0], err
}
