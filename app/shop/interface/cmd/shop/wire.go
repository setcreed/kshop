// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/setcreed/kshop/app/shop/interface/internal/biz"
	"github.com/setcreed/kshop/app/shop/interface/internal/conf"
	"github.com/setcreed/kshop/app/shop/interface/internal/data"
	"github.com/setcreed/kshop/app/shop/interface/internal/server"
	"github.com/setcreed/kshop/app/shop/interface/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
