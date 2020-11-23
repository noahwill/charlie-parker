package helpers

import (
	"charlie-parker/pkg/types"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"
)

// validateCreateRateInput validates a CreateRateInput object and allows for
// optional validation of the inputs against existing rates for overlap
func validateCreateRateInput(in *types.CreateRateInput, checkOverlap bool) error {
	var err error

	if err = validatePrice(in.Price); err != nil {
		return err
	}

	if err = validateTimeZone(in.TZ); err != nil {
		return err
	}

	if err = validateDays(in.Days); err != nil {
		return err
	}

	if err = validateTimespan(in.Times); err != nil {
		return err
	}

	if checkOverlap {
		var rates []types.Rate
		if rates, err = GetRates(); err != nil {
			return err
		}

		if err = validateAgainstExistingRates(rates, *in); err != nil {
			return err
		}
	}

	return err
}

// validateAgainstExistingRates verifies that there is no overlap between new rate being created
// and existing rates
func validateAgainstExistingRates(existingRates []types.Rate, in types.CreateRateInput) error {
	if len(existingRates) > 0 {
		newRanges := getTimeRangesFromDaysAndTimes(in.Days, in.Times, in.TZ, in.Price)
		var existingRanges []timeRange
		for _, existingRate := range existingRates {
			rateRanges := getTimeRangesFromDaysAndTimes(existingRate.Days, existingRate.Times, existingRate.TZ, existingRate.Price)
			existingRanges = append(existingRanges, rateRanges...)
		}

		if err := validateOverlappingRanges(existingRanges, newRanges); err != nil {
			return err
		}
	}

	return nil
}

// validatePrice validates a given price
func validatePrice(price int) error {
	if price == 0 {
		return errors.New("specify a price")
	}

	if price < 0 {
		return errors.New("price must be greater than zero")
	}

	return nil
}

// validateDays validates that days in a comma separated list are valid
// and that there are no repeated days
func validateDays(days string) error {
	if days == "" {
		return errors.New("specify a set of comma separated days")
	}

	daysSlice := strings.Split(days, ",")
	sort.Strings(daysSlice)
	for i, day := range daysSlice {
		if err := isValidDay(day); err != nil {
			return err
		}

		if i+1 < len(daysSlice) {
			dayRemoved := daysSlice[i+1:]
			idx := sort.SearchStrings(dayRemoved, day)
			if dayRemoved[idx] == day {
				return fmt.Errorf("%s is repeated in days", day)
			}
		}
	}
	return nil
}

// validateTimeZone validates a given timezone
func validateTimeZone(tz string) error {
	if tz == "" {
		return errors.New("specify a timezone")
	}

	if _, err := time.LoadLocation(tz); err != nil {
		return fmt.Errorf("invalid timezone: %s", tz)
	}

	return nil
}

// validateTimespan validates a given time range
func validateTimespan(timespan string) error {
	var err error

	if timespan == "" {
		return errors.New("specify a time range")
	}

	var times []string
	if times, err = timeSpanAsSlice(timespan); err != nil {
		return errors.New("specify a time range between only two hours of the day")
	}

	var earlier, later time.Time
	if earlier, later, err = getTimeObjectsFromTimes(times); err != nil {
		return err
	}

	if earlier.After(later) {
		return errors.New("the first time in times must be earlier than the second")
	}

	return err
}

// validateOverlappingRanges errors if there is an overlap between any existingRange
// and any newRange attempting to be created
func validateOverlappingRanges(existingRanges, newRanges []timeRange) error {
	for _, existingRange := range existingRanges {
		for _, newRange := range newRanges {
			// assume that we only care if ranges overlap if they have the same timezone
			if existingRange.tz == newRange.tz {
				overlap := newRange.earlier.Before(existingRange.later) && existingRange.earlier.Before(newRange.later)
				if overlap {
					return fmt.Errorf("a rate already exists for %s %s (TZ: %s, Price: %d) which overlaps the given %s %s (TZ: %s, Price: %d)", existingRange.days, existingRange.times, existingRange.tz, existingRange.price, newRange.days, newRange.times, newRange.tz, newRange.price)
				}
			}
		}
	}
	return nil
}

// validateTimeRange validates that a time range represented by a start and end string
// parses, spans just one year on one day for more than zero seconds, is not mal-ordered,
// and indicates only one timezone
func validateTimeRange(start, end *string) (startTime time.Time, endTime time.Time, err error) {
	if startTime, err = time.Parse(time.RFC3339, *start); err != nil {
		return startTime, endTime, fmt.Errorf("start time parsing error: %v", err)
	}

	if endTime, err = time.Parse(time.RFC3339, *end); err != nil {
		return startTime, endTime, fmt.Errorf("end time parsing error: %v", err)
	}

	if startTime.Year() != endTime.Year() {
		return startTime, endTime, errors.New("start and end cannot be in different years")
	}

	if startTime.YearDay() != endTime.YearDay() {
		return startTime, endTime, errors.New("start and end cannot be on different days")
	}

	if startTime.Equal(endTime) {
		return startTime, endTime, errors.New("start and end cannot be equal")
	}

	if startTime.After(endTime) {
		return startTime, endTime, errors.New("start cannot be after end")
	}

	_, startOffset := startTime.Zone()
	_, endOffset := endTime.Zone()
	if startOffset != endOffset {
		return startTime, endTime, errors.New("start and end cannot be in different timezones")
	}

	return startTime, endTime, err
}
