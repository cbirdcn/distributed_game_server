// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"login_server/internal/biz"
	"login_server/internal/conf"
	"login_server/internal/data"
	"login_server/internal/server"
	"login_server/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	loginServerRepo := data.NewLoginServerRepo(dataData, logger)
	loginServerUsecase := biz.NewLoginServerUsecase(loginServerRepo, logger)
	loginServerService := service.NewLoginServerService(loginServerUsecase)
	grpcServer := server.NewGRPCServer(confServer, loginServerService, logger)
	httpServer := server.NewHTTPServer(confServer, loginServerService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}