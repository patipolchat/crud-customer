package server

import (
	"crud-customer/config"
	"crud-customer/util/server"
	"github.com/labstack/echo/v4"
)

type Server interface {
	SetupServer()
	HttpListening()
	GetEchoApp() *echo.Echo
}

type serverImpl struct {
	server.EchoServer
	AppConfig *config.Config
}

func (s *serverImpl) GetEchoApp() *echo.Echo {
	return s.EchoServer.App
}

func NewServer(cfg *config.Config) Server {
	echoConf := &server.Config{
		Port:         cfg.Server.Port,
		Timeout:      cfg.Server.Timeout,
		AllowOrigins: cfg.Server.AllowOrigins,
		BodyLimit:    cfg.Server.BodyLimit,
		LogLevel:     cfg.Server.LogLevel,
	}
	return &serverImpl{
		EchoServer: server.NewEchoServer(echoConf),
		AppConfig:  cfg,
	}
}
