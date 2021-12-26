//go:build wireinject
// +build wireinject

package repositories

import (
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

func CreateUserRepository(f string) (UserRepository, error) {
	panic(wire.Build(testProviderSet))
}
