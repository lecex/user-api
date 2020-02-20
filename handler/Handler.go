package handler

import (
	"context"

	server "github.com/micro/go-micro/v2/server"
	"github.com/micro/go-micro/v2/util/log"

	authPB "github.com/lecex/user-api/proto/auth"
	casbinPB "github.com/lecex/user-api/proto/casbin"
	frontPermitPB "github.com/lecex/user-api/proto/frontPermit"
	healthPB "github.com/lecex/user-api/proto/health"
	permissionPB "github.com/lecex/user-api/proto/permission"
	rolePB "github.com/lecex/user-api/proto/role"
	userPB "github.com/lecex/user-api/proto/user"
)

// Handler 注册方法
type Handler struct {
	ServiceName string
	Permissions []*permissionPB.Permission
}

// Register 注册
func (srv *Handler) Register(s server.Server) {
	userPB.RegisterUsersHandler(s, &User{srv.ServiceName})
	authPB.RegisterAuthHandler(s, &Auth{srv.ServiceName})
	frontPermitPB.RegisterFrontPermitsHandler(s, &FrontPermit{srv.ServiceName})
	permissionPB.RegisterPermissionsHandler(s, &Permission{srv.ServiceName})
	rolePB.RegisterRolesHandler(s, &Role{srv.ServiceName})
	casbinPB.RegisterCasbinHandler(s, &Casbin{srv.ServiceName}) // 权限管理服务实现
	healthPB.RegisterHealthHandler(s, &Health{})
}

// Sync 同步
func (srv *Handler) Sync() (err error) {
	req := &permissionPB.Request{
		Permissions: srv.Permissions,
	}
	res := &permissionPB.Response{}
	h := Permission{}
	err = h.Sync(context.TODO(), req, res)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
