package http

func (s *Server) setupRoutes() {
	s.server.GET("/healthz", s.handlers.Healthz)

	s.server.GET("/ping", s.handlers.Ping)
}
