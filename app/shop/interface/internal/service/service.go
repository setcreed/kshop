package service

import (
	"github.com/setcreed/kshop/app/shop/interface/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	v1 "github.com/setcreed/kshop/api/shop/interface/v1"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewShopInterface)

type ShopInterface struct {
	v1.UnimplementedShopInterfaceServer

	uc *biz.UserUsecase

	log *log.Helper
}

func NewShopInterface(uc *biz.UserUsecase, logger log.Logger) *ShopInterface {
	return &ShopInterface{
		log: log.NewHelper(log.With(logger, "module", "service/interface")),
		uc:  uc,
	}
}
