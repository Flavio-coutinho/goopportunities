package main

import (
	"github.com/Flavio-coutinho/goopportunities/config"
	"github.com/Flavio-coutinho/goopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}
	// Inicializar router
	router.Initialize()
}
