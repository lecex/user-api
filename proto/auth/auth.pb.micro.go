// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/auth/auth.proto

package auth

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Auth service

type AuthService interface {
	// 用户验证授权
	Auth(ctx context.Context, in *Request, opts ...client.CallOption) (*Request, error)
	// 用户退出登录
	Logout(ctx context.Context, in *Request, opts ...client.CallOption) (*Request, error)
	// token 验证
	ValidateToken(ctx context.Context, in *Request, opts ...client.CallOption) (*Request, error)
}

type authService struct {
	c    client.Client
	name string
}

func NewAuthService(name string, c client.Client) AuthService {
	return &authService{
		c:    c,
		name: name,
	}
}

func (c *authService) Auth(ctx context.Context, in *Request, opts ...client.CallOption) (*Request, error) {
	req := c.c.NewRequest(c.name, "Auth.Auth", in)
	out := new(Request)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) Logout(ctx context.Context, in *Request, opts ...client.CallOption) (*Request, error) {
	req := c.c.NewRequest(c.name, "Auth.Logout", in)
	out := new(Request)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) ValidateToken(ctx context.Context, in *Request, opts ...client.CallOption) (*Request, error) {
	req := c.c.NewRequest(c.name, "Auth.ValidateToken", in)
	out := new(Request)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthHandler interface {
	// 用户验证授权
	Auth(context.Context, *Request, *Request) error
	// 用户退出登录
	Logout(context.Context, *Request, *Request) error
	// token 验证
	ValidateToken(context.Context, *Request, *Request) error
}

func RegisterAuthHandler(s server.Server, hdlr AuthHandler, opts ...server.HandlerOption) error {
	type auth interface {
		Auth(ctx context.Context, in *Request, out *Request) error
		Logout(ctx context.Context, in *Request, out *Request) error
		ValidateToken(ctx context.Context, in *Request, out *Request) error
	}
	type Auth struct {
		auth
	}
	h := &authHandler{hdlr}
	return s.Handle(s.NewHandler(&Auth{h}, opts...))
}

type authHandler struct {
	AuthHandler
}

func (h *authHandler) Auth(ctx context.Context, in *Request, out *Request) error {
	return h.AuthHandler.Auth(ctx, in, out)
}

func (h *authHandler) Logout(ctx context.Context, in *Request, out *Request) error {
	return h.AuthHandler.Logout(ctx, in, out)
}

func (h *authHandler) ValidateToken(ctx context.Context, in *Request, out *Request) error {
	return h.AuthHandler.ValidateToken(ctx, in, out)
}