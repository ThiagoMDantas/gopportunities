package main

import (
	"github.com/ThiagoMDantas/gopportunities/config"
	"github.com/ThiagoMDantas/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err:= config.Init()
	if err != nil {
		logger.Errorf("Config initilization error: %v", err)
		return
	}

	router.Initialize()
}