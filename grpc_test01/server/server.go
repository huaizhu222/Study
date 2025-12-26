package main

import (
	"context"
	"net"

	"Rpc.Study.go/grpc_test01/proto"
	"google.golang.org/grpc"
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
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})

	lis, _ := net.Listen("tcp", "localhost:1234")

	_ = g.Serve(lis)

}
