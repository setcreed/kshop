package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/setcreed/kshop/app/user/service/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) ListUser(ctx context.Context, pageNum, pageSize int) (*[]biz.User, error) {
	panic("implement me")
}
