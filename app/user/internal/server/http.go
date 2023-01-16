package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	api "kratos-realworld/api/user/service/v1"
	"kratos-realworld/app/user/internal/conf"
	"kratos-realworld/app/user/internal/service"
	"kratos-realworld/pkg/middleware/auth"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, jc *conf.Jwt, us *service.UserService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.ErrorEncoder(errorEncoder),
		http.Middleware(
			recovery.Recovery(),
			selector.Server(auth.JwtAuth(jc.Secret)).Match(auth.NewSkipRouterMatcher()).Build(),
			logging.Server(logger),
		),
		http.Filter(
			// 自定义拦截器
			// func(handler nethttp.Handler) nethttp.Handler {
			// 	return nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
			// 		log.Info("route filter in")
			// 		handler.ServeHTTP(w, r)
			// 		log.Info("route filter out")
			// 	})
			// },

			// MDN cors https://developer.mozilla.org/zh-CN/docs/Web/HTTP/CORS
			// OPTIONS请求用于判断是否允许跨域, 两个请求的protocol、port、host都相同的话, 则这两个url是同源
			handlers.CORS(
				handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
				handlers.AllowedOrigins([]string{"*"}),
			)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	api.RegisterUserHTTPServer(srv, us)
	return srv
}
