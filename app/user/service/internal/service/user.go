package service

import (
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/setcreed/kshop/api/user/service/v1"
)

type UserService struct {
	v1.UnimplementedUserServer

	log *log.Helper
}

func NewUserService(logger log.Logger) *UserService {
	return &UserService{
		log: log.NewHelper(log.With(logger, "module", "service/server-service")),
	}
}
