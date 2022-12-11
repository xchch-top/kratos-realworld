package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	realworld "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealworldService)

type RealworldService struct {
	realworld.UnimplementedRealWorldServer

	uc  *biz.UserUseCase
	log *log.Helper
}

func NewRealworldService(uc *biz.UserUseCase, logger log.Logger) *RealworldService {
	return &RealworldService{uc: uc, log: log.NewHelper(logger)}
}
