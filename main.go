package main

import (
	// 公共引入
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/util/log"

	// 执行数据迁移
	"github.com/lecex/user-api/handler"
)

func main() {
	config.LoadFile("config.yaml")
	// 设置权限
	// h := m.Handler{
	// 	Permissions: Conf.Permissions,
	// 	UserService: UserService,
	// }
	service := micro.NewService(
		micro.Name(config.Get("service", "name").String("user-api")),
		micro.Version(config.Get("service", "version").String("latest")),
	)
	service.Init()

	// 注册服务
	handler.Register(service.Server())

	// Run the server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
	log.Log("serviser run ...")
}
