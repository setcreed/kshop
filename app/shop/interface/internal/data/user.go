package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	usv1 "github.com/setcreed/kshop/api/user/service/v1"
	"github.com/setcreed/kshop/app/shop/interface/internal/biz"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/user")),
	}
}

func (rp *userRepo) Register(ctx context.Context, u *biz.User) (*biz.User, error) {
	user, err := rp.data.uc.CreateUser(ctx, &usv1.CreateUserInfo{
		NickName: u.NickName,
		Password: u.Password,
		Mobile:   u.Mobile,
	})
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:     int64(user.Id),
		Mobile: user.Mobile,
	}, nil
}
