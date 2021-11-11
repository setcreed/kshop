package service

import (
	"context"

	v1 "github.com/setcreed/kshop/api/shop/interface/v1"
	"github.com/setcreed/kshop/app/shop/interface/internal/biz"
)

func (s *ShopInterface) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterReply, error) {
	rv, err := s.uc.Register(ctx, &biz.User{
		Mobile:   req.Mobile,
		NickName: req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &v1.RegisterReply{
		Id:     rv.Id,
		Mobile: rv.Mobile,
	}, nil
}
