package biz

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"kratos-realworld/internal/conf"
	"time"
)

type User struct {
	Email        string
	Username     string
	Bio          string
	Image        string
	PassWordHash string
}

type UserLogin struct {
	Email    string
	Username string
	Token    string
	Bio      string
	Image    string
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
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
		PassWordHash: hashPassword(password),
	}
	err := uc.ur.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	return &UserLogin{
		Email:    email,
		Username: username,
		Token:    GenerateToken(uc.jc.Secret, username),
	}, nil
}

func (uc *UserUseCase) Login(ctx context.Context, email string, password string) (*UserLogin, error) {
	u, err := uc.ur.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if !verifyPassword(u.PassWordHash, password) {
		return nil, errors.New(fmt.Sprintf("用户不存在或密码错误"))
	}
	return &UserLogin{
		Email:    u.Email,
		Username: u.Username,
		Bio:      u.Bio,
		Image:    u.Image,
		Token:    GenerateToken(uc.jc.Secret, u.Username),
	}, nil
}

func GenerateToken(secret string, username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"nbf":      time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}

	return tokenString
}
