package service

import (
	"github.com/google/wire"
	realworld "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealworldService)

type RealworldService struct {
	realworld.UnimplementedRealWorldServer

	uc *biz.GreeterUsecase
}

func NewRealworldService(uc *biz.GreeterUsecase) *RealworldService {
	return &RealworldService{uc: uc}
}
