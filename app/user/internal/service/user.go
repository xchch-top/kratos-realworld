package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/golang/protobuf/ptypes/empty"
	api "kratos-realworld/api/user/service/v1"
	"kratos-realworld/app/user/internal/biz"
	"kratos-realworld/pkg/middleware/auth"
)

func (s *UserService) Register(ctx context.Context, req *api.RegisterRequest) (*api.UserLoginReply, error) {
	u, err := s.uc.Register(ctx, req.User.Username, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}

	return &api.UserLoginReply{
		User: &api.UserLoginReply_User{
			Email:    u.Email,
			Username: u.Username,
			Token:    u.Token,
		},
	}, nil
}

func (s *UserService) Login(ctx context.Context, req *api.LoginRequest) (*api.UserLoginReply, error) {
	if req.User.Email == "locked" {
		return nil, errors.InternalServer("email locked", "用户被锁定")
	}

	ul, err := s.uc.Login(ctx, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}

	return &api.UserLoginReply{
		User: &api.UserLoginReply_User{
			Email:    ul.Email,
			Username: ul.Username,
			Bio:      ul.Bio,
			Image:    ul.Image,
			Token:    ul.Token,
		},
	}, nil
}

func (s *UserService) GetCurrentUser(ctx context.Context, empty *empty.Empty) (*api.GetUserReply, error) {
	authUser, err := auth.GetAuthUser(ctx)
	if err != nil {
		return nil, err
	}

	user, err := s.uc.GetUserById(ctx, authUser.Id)
	if err != nil {
		return nil, err
	}

	return &api.GetUserReply{
		User: &api.GetUserReply_User{
			Email:    user.Email,
			Username: user.Username,
			Bio:      user.Bio,
		},
	}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *api.UpdateUserReq) (*api.UserLoginReply, error) {
	authUser, err := auth.GetAuthUser(ctx)
	if err != nil {
		return nil, err
	}

	bizUser := &biz.User{
		Id:       authUser.Id,
		Username: req.Username,
		Email:    req.Email,
		Image:    req.Image,
		Bio:      req.Bio,
	}

	user, err := s.uc.UpdateUser(ctx, bizUser)
	if err != nil {
		return nil, err
	}
	return &api.UserLoginReply{
		User: &api.UserLoginReply_User{
			Username: user.Username,
			Email:    user.Email,
			Token:    user.Token,
			Bio:      user.Bio,
			Image:    user.Image,
		},
	}, nil
}

func (s *UserService) GetProfile(ctx context.Context, req *api.GetProfileRequest) (*api.ProfileReply, error) {
	username := req.GetUsername()
	pf, err := s.uc.GetUserByName(ctx, username)
	if err != nil {
		return nil, err
	}
	return &api.ProfileReply{
		Profile: &api.ProfileReply_Profile{
			Username: pf.Username,
			Bio:      pf.Bio,
			Image:    pf.Image,
		},
	}, nil
}

func (s *UserService) FollowUser(ctx context.Context, req *api.FollowUserRequest) (*api.ProfileReply, error) {
	return nil, nil
}

func (s *UserService) UnfollowUser(ctx context.Context, req *api.UnfollowUserRequest) (*empty.Empty, error) {
	return nil, nil
}
