package handler

import (
	"os"

	server "github.com/micro/go-micro/v2/server"

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
}

// Register 注册
func Register(s server.Server) {
	UserService := os.Getenv("USER_NAME")
	userPB.RegisterUsersHandler(s, &User{UserService})
	authPB.RegisterAuthHandler(s, &Auth{UserService})
	frontPermitPB.RegisterFrontPermitsHandler(s, &FrontPermit{UserService})
	permissionPB.RegisterPermissionsHandler(s, &Permission{UserService})
	rolePB.RegisterRolesHandler(s, &Role{UserService})
	casbinPB.RegisterCasbinHandler(s, &Casbin{UserService}) // 权限管理服务实现
	healthPB.RegisterHealthHandler(s, &Health{})
}
