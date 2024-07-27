package main

import (
	"EffectiveMobile/init/config"
	"EffectiveMobile/init/logger"
	"EffectiveMobile/internal/server"
	"EffectiveMobile/pkg/constants"
	"github.com/sirupsen/logrus"
)

func init() {
	if err := config.InitConfig(); err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.Config})
	}
	logger.Info("Configuration loaded", logrus.Fields{constants.LoggerCategory: constants.Config})
}

// @title Time-Tracker
// @version 1.0.1

// @BasePath /api
func main() {
	app, err := server.NewServer()
	if err != nil {
		logger.Fatal("Error create a new http-server", logrus.Fields{constants.LoggerCategory: constants.Server})
	}
	if err = app.Run(); err != nil {
		logger.Fatal("Error run http-server", logrus.Fields{constants.LoggerCategory: constants.Server})
	}
}
