package main

import (
	"flag"
	"fmt"
	"go-zero-demo/service/search/internal/config"
	"go-zero-demo/service/search/internal/handler"
	"go-zero-demo/service/search/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "search/etc/search-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	//全局中间件
	//server.Use(func(next http.HandlerFunc) http.HandlerFunc {
	//	return func(w http.ResponseWriter, r *http.Request) {
	//		logx.Info("全局中间件")
	//		next(w, r)
	//	}
	//})
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
