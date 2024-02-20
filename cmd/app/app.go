package main

import (
	"github.com/aerosystems/adapter-service/internal/config"
	"github.com/sirupsen/logrus"
	"net/http"
)

type App struct {
	log        *logrus.Logger
	cfg        *config.Config
	httpServer *http.Server
}

func NewApp(log *logrus.Logger, cfg *config.Config, httpServer *http.Server) *App {
	return &App{
		log:        log,
		cfg:        cfg,
		httpServer: httpServer,
	}
}
