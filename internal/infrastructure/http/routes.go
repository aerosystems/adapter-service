package HttpServer

func (s *Server) setupRoutes() {
	s.echo.GET("/domain/check", s.inspectHandler.CheckData, s.AuthTokenMiddleware())
}
