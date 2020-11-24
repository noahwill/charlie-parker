package main

import (
	"os"

	"github.com/labstack/gommon/log"
)

func main() {
	log.Info("This is not the main you are looking for! This main exists to allow for testing.")
	os.Exit(1)
}
