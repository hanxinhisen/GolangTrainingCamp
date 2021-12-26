//go:build wireinject
// +build wireinject

package main

import (
	"blog-users/internal/app/user"
	"blog-users/internal/app/user/controllers"
	"blog-users/internal/app/user/grpcclients"
	"blog-users/internal/app/user/repositories"
	"blog-users/internal/app/user/services"
	"blog-users/internal/pkg/app"
	"blog-users/internal/pkg/config"
	"blog-users/internal/pkg/database"
	"blog-users/internal/pkg/log"
	"blog-users/internal/pkg/transports/grpc"
	"blog-users/internal/pkg/transports/http"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	repositories.ProviderSet,
	services.ProviderSet,
	grpc.ProviderSet,
	grpcclients.ProviderSet,
	http.ProviderSet,
	user.ProviderSet,
	controllers.ProviderSet)

func CreateApp(cf string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
