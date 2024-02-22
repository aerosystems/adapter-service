package HttpServer

import (
	"fmt"
	"github.com/aerosystems/adapter-service/internal/infrastructure/rest"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

const webPort = 80

type Server struct {
	log            *logrus.Logger
	echo           *echo.Echo
	inspectHandler *rest.InspectHandler
	tokenService   TokenService
}

func NewServer(
	log *logrus.Logger,
	inspectHandler *rest.InspectHandler,
	tokenService TokenService,
) *Server {
	return &Server{
		log:            log,
		echo:           echo.New(),
		inspectHandler: inspectHandler,
		tokenService:   tokenService,
	}
}

func (s *Server) Run() error {
	s.setupMiddleware()
	s.setupRoutes()
	s.setupValidator()
	s.log.Infof("starting HTTP server adapter-service on port %d\n", webPort)
	return s.echo.Start(fmt.Sprintf(":%d", webPort))
}
