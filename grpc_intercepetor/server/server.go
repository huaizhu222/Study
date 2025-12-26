package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"Rpc.Study.go/grpc_test01/proto"
	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedGreeterServer // 必须嵌入这个未实现的结构体以符合接口要求
}

func (s *Server) SayHello(c context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	time.Sleep(2 * time.Second)
	return &proto.HelloReply{
		Message: "hello " + request.Name,
	}, nil
}

func main() {
	interceptor := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		fmt.Println("接受了一个新请求")
		return handler(ctx, req)
	}
	opt := grpc.UnaryInterceptor(interceptor)
	g := grpc.NewServer(opt)
	proto.RegisterGreeterServer(g, &Server{})

	lis, _ := net.Listen("tcp", "localhost:1234")

	_ = g.Serve(lis)

}
