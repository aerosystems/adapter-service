//go:build wireinject
// +build wireinject

package main

import (
	"github.com/aerosystems/adapter-service/internal/config"
	HttpServer "github.com/aerosystems/adapter-service/internal/http"
	"github.com/aerosystems/adapter-service/internal/infrastructure/rest"
	"github.com/aerosystems/adapter-service/internal/repository/verifire"
	"github.com/aerosystems/adapter-service/internal/usecases"
	"github.com/aerosystems/adapter-service/pkg/logger"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

//go:generate wire
func InitApp() *App {
	panic(wire.Build(
		wire.Bind(new(rest.ProxyUsecase), new(*usecases.ProxyUsecase)),
		wire.Bind(new(usecases.VerifireRepository), new(*verifire.Api)),
		ProvideApp,
		ProvideLogger,
		ProvideConfig,
		ProvideHttpServer,
		ProvideLogrusLogger,
		ProvideInspectHandler,
		ProvideProxyUsecase,
		ProvideVerifireApi,
	))
}

func ProvideApp(log *logrus.Logger, cfg *config.Config, httpServer *HttpServer.Server) *App {
	panic(wire.Build(NewApp))
}

func ProvideLogger() *logger.Logger {
	panic(wire.Build(logger.NewLogger))
}

func ProvideConfig() *config.Config {
	panic(wire.Build(config.NewConfig))
}

func ProvideHttpServer(log *logrus.Logger, cfg *config.Config, inspectHandler *rest.InspectHandler) *HttpServer.Server {
	return HttpServer.NewServer(log, cfg.AccessSecret, inspectHandler)
}

func ProvideLogrusLogger(log *logger.Logger) *logrus.Logger {
	return log.Logger
}

func ProvideInspectHandler(proxyUsecase rest.ProxyUsecase) *rest.InspectHandler {
	panic(wire.Build(rest.NewInspectHandler))
}

func ProvideProxyUsecase(checkmailRepo usecases.VerifireRepository) *usecases.ProxyUsecase {
	panic(wire.Build(usecases.NewProxyUsecase))
}

func ProvideVerifireApi(cfg *config.Config) *verifire.Api {
	return verifire.NewApi(cfg.VerifireBaseURL)
}
