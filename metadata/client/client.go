package main

import (
	"context"
	"fmt"

	"Rpc.Study.go/grpc_test01/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, _ := grpc.Dial("127.0.0.1:1234", grpc.WithInsecure())
	defer conn.Close()

	c := proto.NewGreeterClient(conn)

	md := metadata.New(map[string]string{
		"name":    "bobby",
		"pasword": "imooc",
	})
	mdd := metadata.Pairs(
		"key1", "val1",
		"key1", "val1-2", // "key1" will have map value []string{"val1", "val1-2"}
		"key2", "val2",
	)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	cttx := metadata.NewOutgoingContext(ctx, mdd)
	r, _ := c.SayHello(cttx, &proto.HelloRequest{
		Name: "hobby",
	})
	fmt.Println(r.Message)
}
