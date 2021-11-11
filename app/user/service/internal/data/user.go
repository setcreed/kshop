package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/setcreed/kshop/app/user/service/internal/biz"
	"github.com/setcreed/kshop/app/user/service/internal/pkg/util"
	"github.com/setcreed/kshop/pkg/util/pagination"
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

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	ph, err := util.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	po, err := r.data.db.User.
		Create().
		SetMobile(u.Mobile).
		SetNickName(u.NickName).
		SetPasswordHash(ph).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:       int32(po.ID),
		Mobile:   po.Mobile,
		NickName: po.NickName,
	}, nil
}

func (r *userRepo) ListUser(ctx context.Context, pageNum, pageSize int64) ([]*biz.User, error) {
	pos, err := r.data.db.User.Query().
		Offset(int(pagination.GetPageOffset(pageNum, pageSize))).
		Limit(int(pageSize)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	rv := make([]*biz.User, 0)
	for _, po := range pos {
		rv = append(rv, &biz.User{
			Id:       int32(po.ID),
			Mobile:   po.Mobile,
			NickName: po.NickName,
			Birthday: po.Birthday.Format("2006-02-01"),
			Gender:   int32(po.Gender),
			Role:     int32(po.Role),
		})
	}
	return rv, err
}
