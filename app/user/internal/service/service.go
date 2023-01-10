package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	user "kratos-realworld/api/user/service/v1"
	"kratos-realworld/app/user/internal/biz"
)

var ProviderSet = wire.NewSet(NewUserService)

type UserService struct {
	user.UnimplementedUserServer

	uc  *biz.UserUseCase
	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/user"))}
}
