package main

import (
	"context"
	"fmt"
	"net"

	"Rpc.Study.go/grpc_test01/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Server struct {
	proto.UnimplementedGreeterServer // 必须嵌入这个未实现的结构体以符合接口要求
}

func (s *Server) SayHello(c context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(c)
	if !ok {
		fmt.Println("error")
	}
	for key, val := range md {
		fmt.Println(key+": ", val)
	}
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
