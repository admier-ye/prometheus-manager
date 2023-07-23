//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"prometheus-manager/apps/master/internal/biz"
	"prometheus-manager/apps/master/internal/conf"
	"prometheus-manager/apps/master/internal/data"
	"prometheus-manager/apps/master/internal/server"
	"prometheus-manager/apps/master/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Bootstrap, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		conf.ProviderSet,
		newApp,
	))
}
