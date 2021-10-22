package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id       int32
	Password string
	Mobile   string
	NickName string
	Birthday string
	Gender   string
	Role     int32
}

type UserRepo interface {
	ListUser(ctx context.Context, pageNum, pageSize int) (*[]User, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/user"))}
}

func (uc *UserUseCase) Get(ctx context.Context, pageNum, pageSize int) (*[]User, error) {
	uc.repo.ListUser(ctx, pageNum, pageSize)
	return nil, nil
}
