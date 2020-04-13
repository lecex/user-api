package handler

import (
	"context"
	"time"

	"github.com/micro/go-micro/v2/util/log"

	server "github.com/micro/go-micro/v2/server"

	client "github.com/lecex/core/client"
	authPB "github.com/lecex/user-api/proto/auth"
	casbinPB "github.com/lecex/user-api/proto/casbin"
	frontPermitPB "github.com/lecex/user-api/proto/frontPermit"
	healthPB "github.com/lecex/user-api/proto/health"
	permissionPB "github.com/lecex/user-api/proto/permission"
	rolePB "github.com/lecex/user-api/proto/role"
	userPB "github.com/lecex/user-api/proto/user"

	"github.com/lecex/user-api/config"
	PB "github.com/lecex/user/proto/permission"
)

// Handler 注册方法
type Handler struct {
	Server server.Server
}

var Conf = config.Conf

// Register 注册
func (srv *Handler) Register() {
	userPB.RegisterUsersHandler(srv.Server, &User{Conf.Service["user"]})
	authPB.RegisterAuthHandler(srv.Server, &Auth{Conf.Service["user"]})
	frontPermitPB.RegisterFrontPermitsHandler(srv.Server, &FrontPermit{Conf.Service["user"]})
	permissionPB.RegisterPermissionsHandler(srv.Server, &Permission{Conf.Service["user"]})
	rolePB.RegisterRolesHandler(srv.Server, &Role{Conf.Service["user"]})
	casbinPB.RegisterCasbinHandler(srv.Server, &Casbin{Conf.Service["user"]}) // 权限管理服务实现
	healthPB.RegisterHealthHandler(srv.Server, &Health{})

	go Sync() // 同步前端权限
}

// Sync 同步
func Sync() {
	time.Sleep(5 * time.Second)
	req := &PB.Request{
		Permissions: Conf.Permissions,
	}
	res := &PB.Response{}
	err := client.Call(context.TODO(), Conf.Service["user"], "Permissions.Sync", req, res)
	if err != nil {
		log.Log(err)
		Sync()
	}
}
