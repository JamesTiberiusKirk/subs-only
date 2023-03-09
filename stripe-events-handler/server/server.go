package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct {
	e *echo.Echo
}

func NewServer(e *echo.Echo) *Server {
	return &Server{
		e: e,
	}
}

func (s *Server) Serve() {
	s.e.GET("/helloworld", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})
}
