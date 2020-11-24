package routes

import (
	"charlie-parker/internal/config"
	"charlie-parker/internal/helpers"
	"charlie-parker/pkg/types"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/gommon/log"

	"github.com/labstack/echo/v4"
)

// GetRatesRoute is the api handler that returns all existing rates from the DB
func GetRatesRoute(c echo.Context) error {
	defer helpers.UpdateRouteResponseTime(time.Now(), helpers.GetRatesRouteName)
	var (
		err        error
		foundRates []types.Rate
		out        types.GetRatesOutput
	)

	if foundRates, err = helpers.GetRates(); err != nil {
		out.Error = fmt.Sprintf("Could not get rates from %s with error: %v", config.Config.RatesTable, err)
		log.Error(out.Error)
		defer helpers.UpdateRouteSuccessFailureCount(false, helpers.GetRatesRouteName)
		return c.JSON(http.StatusInternalServerError, &out)
	}

	out.Ok = true
	out.Rates = foundRates
	log.Infof("Successfully got all %d rates from %s", len(out.Rates), config.Config.RatesTable)
	defer helpers.UpdateRouteSuccessFailureCount(true, helpers.GetRatesRouteName)
	return c.JSON(http.StatusOK, &out)
}

// CreateRateRoute is the api handler for creating a single new rate
func CreateRateRoute(c echo.Context) error {
	defer helpers.UpdateRouteResponseTime(time.Now(), helpers.CreateRateRouteName)
	var (
		err     error
		in      types.CreateRateInput
		newRate types.Rate
		out     types.CreateRateOutput
	)

	if err = c.Bind(&in); err != nil {
		out.Error = fmt.Sprintf("Could not create new rate with error: %v", err)
		log.Error(out.Error)
		defer helpers.UpdateRouteSuccessFailureCount(false, helpers.CreateRateRouteName)
		return c.JSON(http.StatusBadRequest, &out)
	}

	if newRate, err = helpers.CreateRate(&in, true, true); err != nil {
		out.Error = fmt.Sprintf("Could not create rate in %s with error: %v", config.Config.RatesTable, err)
		log.Error(out.Error)
		defer helpers.UpdateRouteSuccessFailureCount(false, helpers.CreateRateRouteName)
		return c.JSON(http.StatusInternalServerError, &out)
	}

	out.Ok = true
	out.Rate = newRate
	log.Infof("Successfully created rate %s rates in %s", out.Rate.UUID, config.Config.RatesTable)
	defer helpers.UpdateRouteSuccessFailureCount(true, helpers.CreateRateRouteName)
	return c.JSON(http.StatusOK, &out)
}

// OverwriteRatesRoute is the api handler that overwrites the existing rates in the DB
func OverwriteRatesRoute(c echo.Context) error {
	defer helpers.UpdateRouteResponseTime(time.Now(), helpers.OverwriteRatesRouteName)
	var (
		err      error
		in       types.OverwriteRatesInput
		newRates []types.Rate
		out      types.OverwriteRatesOutput
	)

	if err = c.Bind(&in); err != nil {
		out.Error = fmt.Sprintf("Could not overwrite rates with error: %v", err)
		log.Error(out.Error)
		defer helpers.UpdateRouteSuccessFailureCount(false, helpers.OverwriteRatesRouteName)
		return c.JSON(http.StatusBadRequest, &out)
	}

	if newRates, err = helpers.OverwriteRates(&in); err != nil {
		out.Error = fmt.Sprintf("Could not overwrite rates in %s with error: %v", config.Config.RatesTable, err)
		log.Error(out.Error)
		defer helpers.UpdateRouteSuccessFailureCount(false, helpers.OverwriteRatesRouteName)
		return c.JSON(http.StatusInternalServerError, &out)
	}

	out.Ok = true
	out.Rates = newRates
	log.Infof("Successfully overwrote existing rates with %d new rates in %s", len(out.Rates), config.Config.RatesTable)
	defer helpers.UpdateRouteSuccessFailureCount(true, helpers.OverwriteRatesRouteName)
	return c.JSON(http.StatusOK, &out)
}

// GetTimespanPriceRoute is the api handler that gets the price from the rate that corresponds to a given date range
func GetTimespanPriceRoute(c echo.Context) error {
	defer helpers.UpdateRouteResponseTime(time.Now(), helpers.GetTimespanPriceRouteName)
	var (
		err   error
		price string
		in    types.GetTimespanPriceInput
		out   types.GetTimespanPriceOutput
	)

	if err = c.Bind(&in); err != nil {
		out.Error = fmt.Sprintf("Could not get price with error: %v", err)
		log.Error(out.Error)
		defer helpers.UpdateRouteSuccessFailureCount(false, helpers.GetTimespanPriceRouteName)
		return c.JSON(http.StatusBadRequest, &out)
	}

	if price, err = helpers.GetTimespanPrice(&in); err != nil {
		out.Error = fmt.Sprintf("Could not get price with error: %v", err)
		log.Error(out.Error)
		defer helpers.UpdateRouteSuccessFailureCount(false, helpers.GetTimespanPriceRouteName)
		return c.JSON(http.StatusInternalServerError, &out)
	}

	out.Ok = true
	out.Price = price
	log.Infof("Successfully got price %s for time range %v -- %v from %s", out.Price, *in.Start, *in.End, config.Config.RatesTable)
	defer helpers.UpdateRouteSuccessFailureCount(true, helpers.GetTimespanPriceRouteName)
	return c.JSON(http.StatusOK, &out)
}
