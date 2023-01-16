package server

import (
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	clientV3 "go.etcd.io/etcd/client/v3"
	"kratos-realworld/app/user/internal/conf"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer, NewRegistrar)

func NewRegistrar(register *conf.Register) registry.Registrar {
	client, err := clientV3.New(clientV3.Config{
		Endpoints: []string{register.Etcd.Endpoints},
	})
	if err != nil {
		panic(err)
	}

	return etcd.New(client)
}
