package HttpServer

import (
	"fmt"
	"github.com/aerosystems/adapter-service/internal/infrastructure/http/handlers"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

const webPort = 80

type Server struct {
	log            *logrus.Logger
	echo           *echo.Echo
	accessSecret   string
	inspectHandler *handlers.InspectHandler
}

func NewServer(
	log *logrus.Logger,
	accessSecret string,
	inspectHandler *handlers.InspectHandler,
) *Server {
	return &Server{
		log:            log,
		echo:           echo.New(),
		accessSecret:   accessSecret,
		inspectHandler: inspectHandler,
	}
}

func (s *Server) Run() error {
	s.setupMiddleware()
	s.setupRoutes()
	s.setupValidator()
	s.log.Infof("starting HTTP server adapter-service on port %d\n", webPort)
	return s.echo.Start(fmt.Sprintf(":%d", webPort))
}
