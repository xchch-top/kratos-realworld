package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
}

type ProfileRepo interface {
}

type UserUseCase struct {
	ur UserRepo
	pr ProfileRepo

	log *log.Helper
}

func NewUserUseCase(ur UserRepo, pr ProfileRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		ur:  ur,
		pr:  pr,
		log: log.NewHelper(logger),
	}
}

func (uc *UserUseCase) Register(ctx context.Context, u *User) error {
	err := uc.ur.CreateUser(ctx, u)
	if err != nil {

	}
	return nil
}
