package service

import (
	"context"
	realworld "kratos-realworld/api/realworld/v1"
)

func (s *RealworldService) Login(ctx context.Context, req *realworld.LoginRequest) (*realworld.UserReply, error) {
	return &realworld.UserReply{
		User: &realworld.User{
			Username: "boom",
		},
	}, nil
}
