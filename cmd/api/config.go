package main

import (
	"github.com/aerosystems/adapter-service/internal/proxy"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Log          *logrus.Logger
	ProxyService proxy.Service
}

func NewApp(
	log *logrus.Logger,
	proxyService *proxy.Service,
) *Config {
	return &Config{
		Log:          log,
		ProxyService: *proxyService,
	}
}
