package main

import (
	// 公共引入
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"

	"github.com/lecex/core/env"
	m "github.com/lecex/user/middleware"

	"github.com/lecex/user-api/config"
	"github.com/lecex/user-api/handler"
)

func main() {
	var Conf = config.Conf
	UserService := env.Getenv("USER_NAME", "user")
	// 设置权限
	h := m.Handler{
		Permissions: Conf.Permissions,
		UserService: UserService,
	}
	service := micro.NewService(
		micro.Name(Conf.Service),
		micro.Version(Conf.Version),
		micro.WrapHandler(h.Wrapper), //验证权限
	)
	service.Init()

	// 注册服务
	handler.Register(service.Server(), UserService)

	// Run the server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
	log.Log("serviser run ...")
}
