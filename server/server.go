package server

import (
	"github.com/labstack/echo/v4"

	"github.com/nogueirahy/api-gymfull/routes"
)

type Server struct {
	port   string
	server *echo.Echo
}

func NewServer() Server {
	return Server{
		port:   ":8000",
		server: echo.New(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	router.Logger.Fatal(router.Start(s.port))
}
