package server

import (
	"crud-customer/config"
	"crud-customer/pkg/echo_server"
	"github.com/labstack/echo/v4"
)

type Server interface {
	SetupServer()
	HttpListening()
	GetEchoApp() *echo.Echo
}

type serverImpl struct {
	echo_server.EchoServer
	AppConfig *config.Config
}

func (s *serverImpl) GetEchoApp() *echo.Echo {
	return s.EchoServer.App
}

func NewServer(cfg *config.Config) Server {
	echoConf := &echo_server.Config{
		Port:         cfg.Server.Port,
		Timeout:      cfg.Server.Timeout,
		AllowOrigins: cfg.Server.AllowOrigins,
		BodyLimit:    cfg.Server.BodyLimit,
		LogLevel:     cfg.Server.LogLevel,
	}
	return &serverImpl{
		EchoServer: echo_server.NewEchoServer(echoConf),
		AppConfig:  cfg,
	}
}
