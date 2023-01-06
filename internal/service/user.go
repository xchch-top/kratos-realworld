package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/golang/protobuf/ptypes/empty"
	realworld "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/biz"
	"kratos-realworld/internal/pkg/middleware/auth"
)

func (s *RealworldService) Register(ctx context.Context, req *realworld.RegisterRequest) (*realworld.UserLoginReply, error) {
	u, err := s.uc.Register(ctx, req.User.Username, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}

	return &realworld.UserLoginReply{
		User: &realworld.UserLoginReply_User{
			Email:    u.Email,
			Username: u.Username,
			Token:    u.Token,
		},
	}, nil
}

func (s *RealworldService) Login(ctx context.Context, req *realworld.LoginRequest) (*realworld.UserLoginReply, error) {
	if req.User.Email == "locked" {
		return nil, errors.InternalServer("email locked", "用户被锁定")
	}

	ul, err := s.uc.Login(ctx, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}

	return &realworld.UserLoginReply{
		User: &realworld.UserLoginReply_User{
			Email:    ul.Email,
			Username: ul.Username,
			Bio:      ul.Bio,
			Image:    ul.Image,
			Token:    ul.Token,
		},
	}, nil
}

func GetAuthUser(ctx context.Context) (*auth.User, error) {
	cVal := ctx.Value(auth.CurrUser)
	if cVal == nil {
		return nil, errors.InternalServer("user not found", "找不到用户")
	}

	cUser := cVal.(auth.User)
	return &cUser, nil
}

func (s *RealworldService) GetCurrentUser(ctx context.Context, empty *empty.Empty) (*realworld.UserReply, error) {
	cUser, err := GetAuthUser(ctx)
	if err != nil {
		return nil, err
	}

	user, err := s.uc.GetCurrent(ctx, cUser.Id)
	if err != nil {
		return nil, err
	}

	return &realworld.UserReply{
		User: &realworld.User{
			Email:    user.Email,
			Username: user.Username,
		},
	}, nil
}

func (s *RealworldService) UpdateUser(ctx context.Context, req *realworld.User) (*realworld.UserLoginReply, error) {
	cUser, err := GetAuthUser(ctx)
	if err != nil {
		return nil, err
	}

	bizUser := &biz.User{Id: cUser.Id, Username: req.Username, Email: req.Email, Image: req.Image, Bio: req.Bio}

	lUser, err := s.uc.UpdateUser(ctx, bizUser)
	if err != nil {
		return nil, err
	}
	return &realworld.UserLoginReply{
		User: &realworld.UserLoginReply_User{
			Username: lUser.Username,
			Email:    lUser.Email,
			Token:    lUser.Token,
			Bio:      lUser.Bio,
			Image:    lUser.Image,
		},
	}, nil
}
