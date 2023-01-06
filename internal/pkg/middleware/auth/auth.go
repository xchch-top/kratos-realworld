package auth

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

var (
	CurrUser = "currUser"
)

func NewSkipRouterMatcher() selector.MatchFunc {
	skipRouters := make(map[string]struct{})
	// /包名.服务名/方法名
	skipRouters["/realworld.v1.RealWorld/Register"] = struct{}{}
	skipRouters["/realworld.v1.RealWorld/Login"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := skipRouters[operation]; ok {
			return false
		}
		return true
	}
}

func JwtAuth(secret string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				tokenString := tr.RequestHeader().Get("Authorization")
				tokenString = strings.SplitN(tokenString, " ", 2)[1]
				log.Info("Authorization: ", tokenString)

				token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
					_, ok := token.Method.(*jwt.SigningMethodHMAC)
					if !ok {
						return nil, errors.Unauthorized("token", "Unexpected signing method: "+token.Header["alg"].(string))
					}
					return []byte(secret), nil
				})
				if err != nil {
					return nil, err
				}

				claims, ok := token.Claims.(jwt.MapClaims)
				if !ok || claims.Valid() != nil || time.Now().Unix()-int64(claims["timestamp"].(float64)) > 30*60 {
					return nil, errors.Unauthorized("token", "token解析失败或token已过期")
				}

				currUser := User{
					Id:       uint64(claims["id"].(float64)),
					Email:    claims["email"].(string),
					Username: claims["username"].(string),
				}
				ctx = context.WithValue(context.Background(), CurrUser, currUser)
				log.Info("username: ", currUser.Email)
			}

			return handler(ctx, req)
		}
	}
}

func GenerateToken(secret string, authUser *User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        authUser.Id,
		"email":     authUser.Email,
		"username":  authUser.Username,
		"timestamp": time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}

	return tokenString
}

type User struct {
	Id       uint64
	Email    string
	Username string
}

func NewAuthUser(id uint64, email string, username string) *User {
	return &User{Id: id, Email: email, Username: username}
}
