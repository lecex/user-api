package config

import (
	m "github.com/lecex/user/middleware"
	PB "github.com/lecex/user/proto/permission"
)

// Config 配置
type Config struct {
	Name        string
	Version     string
	Service     map[string]string
	Permissions []*PB.Permission
}

// Middleware 用户中间件初始化
func (srv *Config) Middleware() *m.Handler {
	return &m.Handler{
		UserService: srv.Service["user"],
		Permissions: srv.Permissions,
	}
}
