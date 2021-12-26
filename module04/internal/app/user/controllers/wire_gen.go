// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package controllers

import (
	"community-blogger/api/protos/user"
	"community-blogger/internal/app/user/repositories"
	"community-blogger/internal/app/user/services"
	"community-blogger/internal/pkg/config"
	"community-blogger/internal/pkg/log"
	"github.com/google/wire"
)

// Injectors from wire.go:

func CreateUserController(cf string, rpo repositories.UserRepository, userclient bloguser.UserClient) (*UserController, error) {
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
	userService := services.NewUserService(logger, viper, rpo, userclient)
	userController := NewUserController(logger, userService)
	return userController, nil
}

// wire.go:

var testProviderSet = wire.NewSet(log.ProviderSet, config.ProviderSet, services.ProviderSet, ProviderSet)