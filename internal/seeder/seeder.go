package seeder

import (
	"charlie-parker/internal/helpers"
	"charlie-parker/pkg/types"
	"os"

	"github.com/labstack/gommon/log"
)

// Run will launch seeding
func Run() {
	log.Infof("Seeding %v Rates", len(rateSeed))
	seed := types.OverwriteRatesInput{Rates: &rateSeed}
	if _, err := helpers.OverwriteRates(&seed); err != nil {
		log.Errorf("Seeding Rates Failed: %v", err)
		os.Exit(1)
	}
	log.Infof("Successfully Seeded Rates")
}
