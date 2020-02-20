package handler

import (
	"context"

	server "github.com/micro/go-micro/v2/server"
	"github.com/micro/go-micro/v2/util/log"

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
	userPB.RegisterUsersHandler(srv.Server, &User{Conf.UserService})
	authPB.RegisterAuthHandler(srv.Server, &Auth{Conf.UserService})
	frontPermitPB.RegisterFrontPermitsHandler(srv.Server, &FrontPermit{Conf.UserService})
	permissionPB.RegisterPermissionsHandler(srv.Server, &Permission{Conf.UserService})
	rolePB.RegisterRolesHandler(srv.Server, &Role{Conf.UserService})
	casbinPB.RegisterCasbinHandler(srv.Server, &Casbin{Conf.UserService}) // 权限管理服务实现
	healthPB.RegisterHealthHandler(srv.Server, &Health{})

	err := srv.Sync() //同步前端数据
	if err != nil {
		log.Fatal(err)
	}
}

// Sync 同步
func (srv *Handler) Sync() (err error) {
	req := &PB.Request{
		Permissions: Conf.Permissions,
	}
	res := &permissionPB.Response{}
	return client.Call(context.TODO(), Conf.UserService, "Permissions.Sync", req, res)
}
