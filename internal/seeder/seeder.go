package seeder

import (
	"github.com/labstack/gommon/log"
)

// Run will launch seeding
func Run() {
	log.Infof("Seeding %v Rates", len(seedData))
}
