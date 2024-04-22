package main

import (
	"atmail/internal/config"
	"atmail/internal/wire"
	"strconv"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// @title          Atmail Assessment Task
// @version        1.0.0
// @description    A Golang HTTP server that performs user management operations (CRUD) using RESTful APIs.
// @contact.name   Jane Marianne Zapanta
// @contact.email  janemarianne.zapanta@gmail.com
// @BasePath       /atmail
// @securityDefinitions.basic BasicAuth
func main() {
	envFlag := config.GetEnvVariable("ENABLE_DEBUG_LOG", "false")
	shouldLogDebug, _ := strconv.ParseBool(envFlag)
	if shouldLogDebug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  config.GetEnvVariable("LOG_INFO_DIR", "info.log"),
		logrus.DebugLevel: config.GetEnvVariable("LOG_DEBUG_DIR", "debug.log"),
		logrus.ErrorLevel: config.GetEnvVariable("LOG_ERROR_DIR", "error.log"),
		logrus.WarnLevel:  config.GetEnvVariable("LOG_WARN_DIR", "warn.log"),
	}

	logrus.AddHook(lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{},
	))

	logrus.SetFormatter(&logrus.JSONFormatter{})

	db, err := config.DB().DB()
	if err != nil {
		logrus.Fatalf(err.Error())
	}
	defer db.Close()

	server := wire.Initialize()
	server.Start()
}
