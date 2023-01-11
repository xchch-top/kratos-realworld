package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/golang/protobuf/ptypes/empty"
	service "kratos-realworld/api/user/service/v1"
	"kratos-realworld/app/user/internal/biz"
	"kratos-realworld/pkg/middleware/auth"
)

func (s *UserService) Register(ctx context.Context, req *service.RegisterRequest) (*service.UserLoginReply, error) {
	u, err := s.uc.Register(ctx, req.User.Username, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}

	return &service.UserLoginReply{
		User: &service.UserLoginReply_User{
			Email:    u.Email,
			Username: u.Username,
			Token:    u.Token,
		},
	}, nil
}

func (s *UserService) Login(ctx context.Context, req *service.LoginRequest) (*service.UserLoginReply, error) {
	if req.User.Email == "locked" {
		return nil, errors.InternalServer("email locked", "用户被锁定")
	}

	ul, err := s.uc.Login(ctx, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}

	return &service.UserLoginReply{
		User: &service.UserLoginReply_User{
			Email:    ul.Email,
			Username: ul.Username,
			Bio:      ul.Bio,
			Image:    ul.Image,
			Token:    ul.Token,
		},
	}, nil
}

func (s *UserService) GetCurrentUser(ctx context.Context, empty *empty.Empty) (*service.GetUserReply, error) {
	authUser, err := auth.GetAuthUser(ctx)
	if err != nil {
		return nil, err
	}

	user, err := s.uc.GetCurrent(ctx, authUser.Id)
	if err != nil {
		return nil, err
	}

	return &service.GetUserReply{
		User: &service.GetUserReply_User{
			Email:    user.Email,
			Username: user.Username,
		},
	}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *service.UpdateUserReq) (*service.UserLoginReply, error) {
	authUser, err := auth.GetAuthUser(ctx)
	if err != nil {
		return nil, err
	}

	bizUser := &biz.User{Id: authUser.Id, Username: req.Username, Email: req.Email, Image: req.Image, Bio: req.Bio}

	lUser, err := s.uc.UpdateUser(ctx, bizUser)
	if err != nil {
		return nil, err
	}
	return &service.UserLoginReply{
		User: &service.UserLoginReply_User{
			Username: lUser.Username,
			Email:    lUser.Email,
			Token:    lUser.Token,
			Bio:      lUser.Bio,
			Image:    lUser.Image,
		},
	}, nil
}

func (s *UserService) GetProfile(ctx context.Context, req *service.GetProfileRequest) (*service.ProfileReply, error) {
	return nil, nil
}

func (s *UserService) FollowUser(ctx context.Context, req *service.FollowUserRequest) (*service.ProfileReply, error) {
	return nil, nil
}

func (s *UserService) UnfollowUser(ctx context.Context, req *service.UnfollowUserRequest) (*empty.Empty, error) {
	return nil, nil
}
