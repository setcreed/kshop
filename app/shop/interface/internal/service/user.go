package service

import (
	"context"

	v1 "github.com/setcreed/kshop/api/shop/interface/v1"
)

func (s ShopInterface) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterReply, error) {
	panic("implement me")
}
