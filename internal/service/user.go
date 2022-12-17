package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	realworld "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/errors"
)

func (s *RealworldService) Register(ctx context.Context, req *realworld.RegisterRequest) (*realworld.UserReply, error) {
	u, err := s.uc.Register(ctx, req.User.Username, req.User.Username, req.User.Password)
	if err != nil {
		return nil, err
	}

	return &realworld.UserReply{
		User: &realworld.User{
			Username: u.Username,
			Token:    u.Token,
		},
	}, nil
}

func (s *RealworldService) Login(ctx context.Context, req *realworld.LoginRequest) (*realworld.UserReply, error) {
	if req.User.Email == "locked" {
		// 错误时返回用例: {"errors":{"email":"用户被锁定"}}
		return nil, errors.NewHttpError(422, "email", "用户被锁定")
	}

	return &realworld.UserReply{
		User: &realworld.User{
			Username: "boom",
		},
	}, nil
}

func (s *RealworldService) GetCurrentUser(ctx context.Context, empty *empty.Empty) (*realworld.UserReply, error) {
	return nil, nil
}

func (s *RealworldService) UpdateUser(ctx context.Context, req *realworld.UpdateUserRequest) (*realworld.UserReply, error) {
	return nil, nil
}
