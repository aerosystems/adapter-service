package main

import (
	"fmt"
	"github.com/aerosystems/adapter-service/internal/proxy"
	"github.com/aerosystems/adapter-service/pkg/logger"
	"os"
)

const webPort = "80"

// @title Adapter Service API
// @version 1.0.0
// @description A part of microservice infrastructure, who responsible for proxy requests to checkmail-service

// @contact.name Artem Kostenko
// @contact.url https://github.com/aerosystems

// @license.name Apache 2.0
// @license.url https://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Should contain Access JWT Token, with the Bearer started

// @host api.verifire.app
// @schemes https
// @BasePath /
func main() {
	log := logger.NewLogger(os.Getenv("HOSTNAME"))
	proxyService := proxy.NewService("http://checkmail-service:80")
	app := NewApp(log.Logger, proxyService)

	e := app.NewRouter()

	err := e.Start(fmt.Sprintf(":%s", webPort))
	if err != nil {
		panic(err)
	}
}
