package handler

import (
	"context"
	"fmt"
	"time"

	client "github.com/lecex/core/client"
	"github.com/micro/go-micro/v2/util/log"

	pb "github.com/lecex/user-api/proto/auth"
	authSrvPB "github.com/lecex/user/proto/auth"

	"github.com/go-redis/redis"
)

// Auth 授权服务处理
type Auth struct {
	ServiceName string
	Redis       *redis.Client
}

// Auth 授权认证
// 返回token
func (srv *Auth) Auth(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	err = client.Call(ctx, srv.ServiceName, "Auth.Auth", req, res)
	return err
}

// Mobile 手机验证码授权
func (srv *Auth) Mobile(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// if req.User.Mobile != "" { // 验证手机
	// 	err = srv.VerifyCaptcha(req.User.Mobile, req.Captcha)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	reqAuthSrv := &authSrvPB.Request{
		User: &authSrvPB.User{
			Mobile: req.User.Mobile,
		},
	}
	log.Log(req)
	log.Log(reqAuthSrv)
	resAuthSrv := &authSrvPB.Response{}
	err = client.Call(ctx, srv.ServiceName, "Auth.AuthById", reqAuthSrv, resAuthSrv)
	if err != nil {
		return err
	}
	log.Log(resAuthSrv)
	res.Token = resAuthSrv.Token
	return err
}

// Logout 登录退出
// 后期开发
// token 增加前期认证根据根据存储的 token 进行认证 和过期查询
// 通过删除 存储 token 来登出
// 通过存储多种类型的 token 来实现多端登录
// token id type type=登录类型
// 过期时间默认还是在 jwt token 中存储
func (srv *Auth) Logout(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return err
}

// ValidateToken 验证 token
// 并且验证 token 所属用户相关权限
func (srv *Auth) ValidateToken(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Auth.ValidateToken", req, res)
}

// VerifyCaptcha 校验验证码
func (srv *Auth) VerifyCaptcha(addressee string, captcha string) (err error) {
	r, err := srv.Redis.Get("Captcha_" + addressee).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			return fmt.Errorf("验证码已超时")
		}
		return err
	}
	if r != captcha {
		return fmt.Errorf("验证码错误")
	}
	srv.Redis.Set("Captcha_"+addressee, r, 1*time.Second).Err() // 1秒后自动过期
	return nil

}
