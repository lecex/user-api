package config

import (
	m "github.com/lecex/user/middleware"
	PB "github.com/lecex/user/proto/permission"
)

// Config 配置
type Config struct {
	Service     string
	Version     string
	UserService string
	Permissions []*PB.Permission
}

// Middleware 用户中间件初始化
func (srv *Config) Middleware() *m.Handler {
	return &m.Handler{
		UserService: srv.UserService,
		Permissions: srv.Permissions,
	}
}
