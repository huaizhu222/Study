package main

import (
	"context"
	"fmt"

	"Rpc.Study.go/grpc_test01/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial("127.0.0.1:1234", grpc.WithInsecure())
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	r, _ := c.SayHello(context.Background(), &proto.HelloRequest{
		Name: "hobby",
	})
	fmt.Println(r.Message)
}
