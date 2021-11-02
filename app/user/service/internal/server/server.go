package server

import (
	"github.com/go-kratos/kratos/v2/registry"
	nacos "github.com/go-kratos/nacos/registry"
	"github.com/google/wire"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"

	"github.com/setcreed/kshop/app/user/service/internal/conf"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewRegistrar)

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(conf.Nacos.IpAddr, conf.Nacos.Port),
	}

	cc := &constant.ClientConfig{
		NamespaceId:         conf.Nacos.NamespaceID,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              conf.Nacos.LogDir,
		CacheDir:            conf.Nacos.CacheDir,
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}
	cli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}

	return nacos.New(cli)

}
