package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id       int32
	Mobile   string
	Password string
	NickName string
	Birthday string
	Gender   int32
	Role     int32
}

type UserRepo interface {
	CreateUser(ctx context.Context, u *User) (*User, error)
	ListUser(ctx context.Context, pageNum, pageSize int64) ([]*User, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/user"))}
}

func (uc *UserUseCase) Get(ctx context.Context, pageNum, pageSize int64) ([]*User, error) {
	return uc.repo.ListUser(ctx, pageNum, pageSize)
}

func (uc *UserUseCase) Create(ctx context.Context, u *User) (*User, error) {
	out, err := uc.repo.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}
	return out, nil
}
