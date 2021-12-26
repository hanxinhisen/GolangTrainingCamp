// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//go:build !wireinject
// +build !wireinject

package controllers

import (
	"blog-users/api/protos/user"
	"blog-users/internal/app/user/repositories"
	"blog-users/internal/app/user/services"
	"blog-users/internal/pkg/config"
	"blog-users/internal/pkg/log"
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
