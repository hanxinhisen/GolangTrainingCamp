// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"community-blogger/internal/app/user"
	"community-blogger/internal/app/user/controllers"
	"community-blogger/internal/app/user/grpcclients"
	"community-blogger/internal/app/user/repositories"
	"community-blogger/internal/app/user/services"
	"community-blogger/internal/pkg/app"
	"community-blogger/internal/pkg/config"
	"community-blogger/internal/pkg/database"
	"community-blogger/internal/pkg/log"
	"community-blogger/internal/pkg/transports/grpc"
	"community-blogger/internal/pkg/transports/http"
	"github.com/google/wire"
)

// Injectors from wire.go:

func CreateApp(cf string) (*app.Application, error) {
	viper, err := config.New(cf)
	if err != nil {
		return nil, err
	}
	options, err := log.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	logger, err := log.New(options)
	if err != nil {
		return nil, err
	}
	userOptions, err := user.NewOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	httpOptions, err := http.NewOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	databaseOptions, err := database.NewOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	databaseDatabase, err := database.New(databaseOptions)
	if err != nil {
		return nil, err
	}
	userRepository := repositories.NewMysqlUserRepository(logger, databaseDatabase)
	clientOptions, err := grpc.NewClientOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	client, err := grpc.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}
	userClient, err := grpcclients.NewUserClient(client, viper)
	if err != nil {
		return nil, err
	}
	userService := services.NewUserService(logger, viper, userRepository, userClient)
	userController := controllers.NewUserController(logger, userService)
	initControllers := controllers.CreateInitControllersFn(userController)
	engine := http.NewRouter(httpOptions, logger, initControllers)
	server, err := http.New(httpOptions, logger, engine)
	if err != nil {
		return nil, err
	}
	application, err := user.NewApp(userOptions, logger, server)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(log.ProviderSet, config.ProviderSet, database.ProviderSet, repositories.ProviderSet, services.ProviderSet, grpc.ProviderSet, grpcclients.ProviderSet, http.ProviderSet, user.ProviderSet, controllers.ProviderSet)