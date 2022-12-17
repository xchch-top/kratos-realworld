package auth

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
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
						return nil, errors.New("Unexpected signing method: " + token.Header["alg"].(string))
					}
					return []byte(secret), nil
				})
				if err != nil {
					return nil, err
				}

				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					log.Info("username: ", claims["username"])
				} else {
					return nil, err
				}
			}

			return handler(ctx, req)
		}
	}
}

func GenerateToken(secret string, username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  username,
		"timestamp": time.Now().UTC(),
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}

	return tokenString
}
