package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

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

	srv.server.GET("/api/user", srv.GetUser)
	srv.server.POST("/api/user", srv.AddUser)

	return srv
}

func (srv *Server) Run() {
	srv.server.Logger.Fatal(srv.server.Start(srv.address))
}
