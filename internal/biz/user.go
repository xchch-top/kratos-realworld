package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"kratos-realworld/internal/conf"
	"kratos-realworld/internal/pkg/middleware/auth"
)

type User struct {
	Id           uint64
	Email        string
	Username     string
	Bio          string
	Image        string
	PasswordHash string
}

type UserLogin struct {
	Email    string
	Username string
	Token    string
	Bio      string
	Image    string
}

type UserRepo interface {
	CreateUser(ctx context.Context, u *User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserById(ctx context.Context, id uint64) (*User, error)
	UpdateUser(ctx context.Context, user *User) (*User, error)
}

func hashPassword(pwd string) string {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	log.Info("password", pwd, ", hashedPwd", string(hashedPwd))
	return string(hashedPwd)
}

func verifyPassword(hashedPwd string, input string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(input))
	return err == nil
}

type ProfileRepo interface {
}

type UserUseCase struct {
	ur UserRepo
	pr ProfileRepo

	log *log.Helper
	jc  *conf.Jwt
}

func NewUserUseCase(ur UserRepo, pr ProfileRepo, logger log.Logger, jc *conf.Jwt) *UserUseCase {
	return &UserUseCase{
		ur:  ur,
		pr:  pr,
		log: log.NewHelper(logger),
		jc:  jc,
	}
}

func (uc *UserUseCase) Register(ctx context.Context, username string, email string, password string) (*UserLogin, error) {
	u := &User{
		Email:        email,
		Username:     username,
		PasswordHash: hashPassword(password),
	}

	// 判断邮箱是否已经注册过
	uByEmail, err := uc.ur.GetUserByEmail(ctx, email)
	if uByEmail != nil {
		return nil, errors.InternalServer("email has bean registered", "邮箱已经被注册过")
	}

	// 注册用户
	bizUser, err := uc.ur.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	return &UserLogin{
		Email:    email,
		Username: username,
		Token:    auth.GenerateToken(uc.jc.Secret, auth.NewAuthUser(bizUser.Id, bizUser.Email, bizUser.Username)),
	}, nil
}

func (uc *UserUseCase) Login(ctx context.Context, email string, password string) (*UserLogin, error) {
	u, err := uc.ur.GetUserByEmail(ctx, email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.InternalServer("user or password error", "用户不存在或密码错误")
	}
	if err != nil {
		return nil, err
	}

	if !verifyPassword(u.PasswordHash, password) {
		// 项目中统一使用kratos的error包
		return nil, errors.InternalServer("user or password error", "用户不存在或密码错误")
	}
	return &UserLogin{
		Email:    u.Email,
		Username: u.Username,
		Bio:      u.Bio,
		Image:    u.Image,
		Token:    auth.GenerateToken(uc.jc.Secret, auth.NewAuthUser(u.Id, u.Email, u.Username)),
	}, nil
}

func (uc *UserUseCase) GetCurrent(ctx context.Context, id uint64) (*User, error) {
	u, err := uc.ur.GetUserById(ctx, id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		errors.NotFound("user not found", "用户不存在")
	}
	if err != nil {
		return nil, err
	}

	return u, err
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, newUser *User) (*UserLogin, error) {
	oUser, err := uc.ur.GetUserById(ctx, newUser.Id)

	if errors.Is(err, gorm.ErrRecordNotFound) || oUser == nil {
		errors.NotFound("user not found", "用户不存在")
	}
	if err != nil {
		return nil, err
	}

	// todo 邮箱
	// todo 如果密码字段不为空则改密码

	newUser, err = uc.ur.UpdateUser(ctx, newUser)
	if err != nil {
		errors.InternalServer("user update failed", "用户信息更新失败")
	}

	return &UserLogin{
		Email:    newUser.Email,
		Username: newUser.Username,
		Bio:      newUser.Bio,
		Image:    newUser.Image,
		Token:    auth.GenerateToken(uc.jc.Secret, auth.NewAuthUser(newUser.Id, newUser.Email, newUser.Username)),
	}, nil
}
