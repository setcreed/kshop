package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	v1 "github.com/setcreed/kshop/api/shop/interface/v1"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewShopInterface)

type ShopInterface struct {
	v1.UnimplementedShopInterfaceServer

	log *log.Helper
}

func (s ShopInterface) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterReply, error) {
	panic("implement me")
}

func NewShopInterface(logger log.Logger) *ShopInterface {
	return &ShopInterface{
		log: log.NewHelper(log.With(logger, "module", "service/interface")),
	}
}
