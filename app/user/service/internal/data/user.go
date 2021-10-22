package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/setcreed/kshop/pkg/util/pagination"

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
			Birthday: "",
			Gender:   int32(po.Gender),
			Role:     int32(po.Role),
		})
	}
	return rv, err
}
