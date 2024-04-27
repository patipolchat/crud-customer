package app

import (
	"crud-customer/config"
	"crud-customer/database"
	v1 "crud-customer/routes/api/v1"
	"crud-customer/server"
)

type App struct {
	Config *config.Config
	Server server.Server
	DB     database.GormDB
}

func NewApp(cfg *config.Config, db database.GormDB, server server.Server) *App {
	return &App{
		Config: cfg,
		DB:     db,
		Server: server,
	}
}

func (a *App) Start() {
	a.Server.SetupServer()
	a.SetupRoute()
	a.Server.HttpListening()
}

func (a *App) SetupRoute() {
	v1.SetCustomerRoutes(a.Config, a.Server.GetEchoApp(), a.DB)
}
