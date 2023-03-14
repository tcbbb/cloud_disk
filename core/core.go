package main

import (
	"flag"
	"fmt"

	"cloud-disk/core/internal/config"
	"cloud-disk/core/internal/handler"
	"cloud-disk/core/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/core-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	// 将yaml文件配置导入config对象
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()
	// 注册路由
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
