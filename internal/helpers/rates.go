package helpers

import (
	"charlie-parker/internal/config"
	"charlie-parker/pkg/types"
	"errors"
	"strconv"
	"time"

	"github.com/gofrs/uuid"
	"github.com/labstack/gommon/log"
)

// GetRates gets all of the rates from the DB
func GetRates() ([]types.Rate, error) {
	var rates []types.Rate
	err := config.Config.RatesTableConn.Scan().Filter("$ = ?", "Active", true).All(&rates)
	return rates, err
}

// CreateRate creates a rate in the DB and allows for optional validation of the inputs
// against existing rates for overlap and the option to create the rate in the DB immediately
func CreateRate(in *types.CreateRateInput, checkOverlap bool, createImmediately bool) (types.Rate, error) {
	var (
		err  error
		rate types.Rate
	)

	if err = validateCreateRateInput(in, checkOverlap); err != nil {
		return rate, err
	}

	uu, _ := uuid.NewV4()
	log.Infof("UUID: %s", uu.String())
	rate = types.Rate{
		Days:  in.Days,
		Times: in.Times,
		TZ:    in.TZ,
		Price: in.Price,
		UUID:  uu.String(),
	}

	if createImmediately {
		if err = putRatesInTable(rate); err != nil {
			return rate, err
		}
	}

	return rate, err
}

// putRatesInTable puts one or more rates in the RatesTable
func putRatesInTable(rates ...types.Rate) error {
	for _, rate := range rates {
		if err := config.Config.RatesTableConn.Put(&rate).Run(); err != nil {
			return err
		}
	}
	return nil
}

// OverwriteRates deletes all existing rates and replaces them with new ones from input
func OverwriteRates(in *types.OverwriteRatesInput) ([]types.Rate, error) {
	var (
		err   error
		rates []types.Rate
	)

	if in.Rates == nil {
		return rates, errors.New("specify at least 1 rate to create")
	} else if len(*in.Rates) == 0 {
		return rates, errors.New("specify at least 1 rate to create")
	}

	for _, input := range *in.Rates {
		var rate types.Rate
		if rate, err = CreateRate(&input, true, false); err != nil {
			return rates, err
		}
		rates = append(rates, rate)
	}

	var oldRates []types.Rate
	if oldRates, err = GetRates(); err != nil {
		return rates, err
	}

	for _, oldRate := range oldRates {
		if err = config.Config.RatesTableConn.Delete("UUID", oldRate.UUID).Run(); err != nil {
			return rates, err
		}
	}
	err = putRatesInTable(rates...)
	return rates, err
}

// GetTimespanPrice finds the price corresponding to the given input
func GetTimespanPrice(in *types.GetTimespanPriceInput) (string, error) {
	var (
		err                error
		price              string = "unavailable"
		matchedRate        types.Rate
		startTime, endTime time.Time
	)

	if in.Start == nil {
		return price, errors.New("specify start")
	} else if in.End == nil {
		return price, errors.New("specify end")
	}

	if startTime, endTime, err = validateTimeRange(in.Start, in.End); err != nil {
		return price, err
	}

	if matchedRate, err = matchTimespanToRate(startTime, endTime); err != nil {
		return price, err
	}

	price = strconv.Itoa(matchedRate.Price)
	return price, err
}
