package service

import (
	"context"
	pb "github.com/sexiong306/aibum-protocol/helloworld"
)

type GreeterImpl struct {
}

// Hello 函数入口，用户逻辑写在该函数内部即可
// error 代表的是 exception，异常情况比如数据库连接错误，调用下游服务错误的时候，如果返回 error，rsp 的内容将不再被返回
// 如果业务遇到需要返回的错误码，错误信息，而且同时需要保留 HelloReply，请设计在 HelloReply 里面，并将 error 返回 nil
func (s *GreeterImpl) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	rsp := &pb.HelloReply{}
	rsp.Msg = "accept"
	return rsp, nil
}
