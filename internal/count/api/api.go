package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Usecase interface {
	FetchCount() (int, error)
	IncreaseCount(int) error
}

type Server struct {
	server  *echo.Echo
	address string
	uc      Usecase
}

func NewServer(ip string, port int, uc Usecase) *Server {
	srv := &Server{
		server:  echo.New(),
		address: fmt.Sprintf("%s:%d", ip, port),
		uc:      uc,
	}

	srv.server.GET("/count", srv.GetCount)
	srv.server.POST("/count", srv.IncreaseCount)

	return srv
}

func (srv *Server) Run() {
	srv.server.Logger.Fatal(srv.server.Start(srv.address))
}
