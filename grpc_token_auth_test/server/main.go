package main

import (
	"context"
	"fmt"
	"net"

	"Rpc.Study.go/grpc_test01/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Server struct {
	proto.UnimplementedGreeterServer // 必须嵌入这个未实现的结构体以符合接口要求
}

func (s *Server) SayHello(c context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "hello " + request.Name,
	}, nil
}

func main() {
	interceptor := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		fmt.Println("接受了一个新请求")
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return resp, status.Error(codes.Unauthenticated, "无token认证信息")
		}
		for key, val := range md {
			fmt.Println(key+": ", val)
		}
		return handler(ctx, req)
	}
	opt := grpc.UnaryInterceptor(interceptor)
	g := grpc.NewServer(opt)
	proto.RegisterGreeterServer(g, &Server{})

	lis, _ := net.Listen("tcp", "localhost:1234")

	_ = g.Serve(lis)

}
