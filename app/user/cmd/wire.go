//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"kratos-realworld/app/user/internal/biz"
	"kratos-realworld/app/user/internal/conf"
	"kratos-realworld/app/user/internal/data"
	"kratos-realworld/app/user/internal/server"
	"kratos-realworld/app/user/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Jwt, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
