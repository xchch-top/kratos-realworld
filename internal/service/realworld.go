package service

import (
	"context"

	realworld "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/biz"
)

type RealworldService struct {
	realworld.UnimplementedRealWorldServer

	uc *biz.GreeterUsecase
}

func NewRealworldService(uc *biz.GreeterUsecase) *RealworldService {
	return &RealworldService{uc: uc}
}

func (s *RealworldService) SayHello(ctx context.Context, in *realworld.HelloRequest) (*realworld.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &realworld.HelloReply{Message: "Hello " + g.Hello}, nil
}
