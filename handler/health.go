package handler

import (
	"context"

	pb "github.com/lecex/user-api/proto/health"
)

// Health 用户结构
type Health struct {
}

// Health 用户是否存在
func (srv *Health) Health(ctx context.Context, req *pb.Request, rsp *pb.Response) (err error) {
	rsp.StatusCode = 200
	rsp.Body = string("true")
	return nil
}
