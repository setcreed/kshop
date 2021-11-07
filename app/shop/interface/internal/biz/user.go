package biz

import "github.com/go-kratos/kratos/v2/log"

type UserRepo interface {
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{
		log:  log.NewHelper(log.With(logger, "module", "usecase/interface")),
		repo: repo,
	}
}
