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

func (s *UserService) GetUserList(ctx context.Context, req *v1.PageInfo) (*v1.UserListResponse, error) {
	return nil, nil
}
