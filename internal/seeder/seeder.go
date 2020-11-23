package seeder

import (
	"charlie-parker/internal/helpers"
	"charlie-parker/pkg/types"

	"github.com/labstack/gommon/log"
)

// Run will launch seeding
func Run() {
	log.Infof("Seeding %v Rates", len(seedData))
	seed := types.OverwriteRatesInput{Rates: &seedData}
	if _, err := helpers.OverwriteRates(&seed); err != nil {
		log.Infof("Seeding failed: %v", err)
	}
}
