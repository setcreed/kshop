package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	v1 "github.com/setcreed/kshop/api/user/service/v1"
	"github.com/setcreed/kshop/app/user/service/internal/biz"
)

type UserService struct {
	v1.UnimplementedUserServer

	uc  *biz.UserUseCase
	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/server-service")),
	}
}

func (s *UserService) GetUserList(ctx context.Context, req *v1.PageInfo) (*v1.UserResponseList, error) {
	userList, err := s.uc.Get(ctx, int64(req.PageNum), int64(req.PageSize))
	if err != nil {
		return nil, err
	}

	rs := make([]*v1.UserInfoResponse, 0)

	for _, info := range userList {
		rs = append(rs, &v1.UserInfoResponse{
			Id:       info.Id,
			Mobile:   info.Mobile,
			NickName: info.NickName,
			Birthday: info.Birthday,
			Gender:   info.Gender,
			Role:     info.Role,
		})
	}

	return &v1.UserResponseList{
		Total: int32(len(rs)),
		Data:  rs,
	}, nil
}
