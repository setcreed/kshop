package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id       int64
	Mobile   string
	NickName string
	Password string
}

type UserRepo interface {
	Register(ctx context.Context, u *User) (*User, error)
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

func (uc *UserUsecase) Register(ctx context.Context, u *User) (*User, error) {
	return uc.repo.Register(ctx, u)
}
