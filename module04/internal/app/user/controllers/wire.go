//go:build wireinject
// +build wireinject

package controllers

import (
	userproto "blog-users/api/protos/user"
	"blog-users/internal/app/user/repositories"
	"blog-users/internal/app/user/services"
	"blog-users/internal/pkg/config"
	"blog-users/internal/pkg/log"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	services.ProviderSet,
	ProviderSet)

func CreateUserController(cf string,
	rpo repositories.UserRepository,
	userclient userproto.UserClient) (*UserController, error) {
	panic(wire.Build(testProviderSet))
}
