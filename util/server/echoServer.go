package server

import (
	"context"
	"crud-customer/util/validator"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Config struct {
	Port         int
	Timeout      time.Duration
	AllowOrigins []string
	BodyLimit    string
	LogLevel     string
}

type EchoServer struct {
	App        *echo.Echo
	EchoConfig *Config
}

func (s *EchoServer) GetLogLevel() log.Lvl {
	switch s.EchoConfig.LogLevel {
	case "DEBUG":
		return log.DEBUG
	case "INFO":
		return log.INFO
	case "WARN":
		return log.WARN
	case "ERROR":
		return log.ERROR
	default:
		fmt.Printf("Invalid log level: %s. Defaulting to DEBUG", s.EchoConfig.LogLevel)
		return log.DEBUG
	}
}

func (s *EchoServer) GetRouter() *echo.Router {
	return s.App.Router()
}

func (s *EchoServer) SetupServer() {
	s.App.Logger.SetLevel(s.GetLogLevel())
	s.App.Validator = validator.GetEchoValidator()
	s.setupMiddleWares()
}

func (s *EchoServer) setupMiddleWares() {
	s.App.Use(middleware.Recover())
	s.App.Use(middleware.Logger())
	s.App.Use(GetTimeOutMiddleware(s.EchoConfig.Timeout))
	s.App.Use(GetCORSMiddleware(s.EchoConfig.AllowOrigins))
	s.App.Use(GetBodyLimitMiddleware(s.EchoConfig.BodyLimit))
}

func (s *EchoServer) HttpListening() {
	url := fmt.Sprintf(":%d", s.EchoConfig.Port)
	s.setupGracefullyShutdown()
	if err := s.App.Start(url); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.App.Logger.Panicf("Error: %v", err)
	}
}

func (s *EchoServer) setupGracefullyShutdown() {
	ctx := context.Background()
	quitCh := make(chan os.Signal, 1)
	go func(quitCh chan os.Signal) {
		signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)
		<-quitCh
		s.App.Logger.Infof("Shutting down service...")

		if err := s.App.Shutdown(ctx); err != nil {
			s.App.Logger.Fatalf("Error: %s", err.Error())
		}
	}(quitCh)
}

func NewEchoServer(cfg *Config) EchoServer {
	echoApp := echo.New()

	return EchoServer{
		App:        echoApp,
		EchoConfig: cfg,
	}
}
