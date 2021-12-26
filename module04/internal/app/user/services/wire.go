//go:build wireinject
// +build wireinject

package services

import (
	userproto "blog-users/api/protos/user"
	"blog-users/internal/app/user/repositories"
	"blog-users/internal/pkg/config"
	"blog-users/internal/pkg/database"
	"blog-users/internal/pkg/log"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	ProviderSet)

func CreateUserService(cf string,
	rpo repositories.UserRepository,
	userclient userproto.UserClient) (UserService, error) {
	panic(wire.Build(testProviderSet))
}
