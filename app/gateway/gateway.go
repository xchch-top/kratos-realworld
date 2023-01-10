package main

import (
	"github.com/gin-gonic/gin"
	"net/http/httputil"
	"net/url"
	"strings"
)

// 通过gin实现一个反向代理服务器 https://zhuanlan.zhihu.com/p/576833062
func main() {
	router := gin.Default()

	router.Any("/api/*params", func(ctx *gin.Context) {
		host := "localhost:8001"
		if strings.Contains(ctx.Request.RequestURI, "/api/article") {
			host = "localhost:8002"
		}

		remote, err := url.Parse("http://" + host)
		if err != nil {
			panic(err)
		}
		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.ServeHTTP(ctx.Writer, ctx.Request)
	})

	router.Run(":8000")
}
