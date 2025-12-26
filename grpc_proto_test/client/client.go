package main

import (
	"context"
	"fmt"
	"time"

	proto_bak "Rpc.Study.go/grpc_proto_test/proto-bak"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	conn, _ := grpc.Dial("127.0.0.1:1234", grpc.WithInsecure())
	defer conn.Close()
	c := proto_bak.NewGreeterClient(conn)
	r, _ := c.SayHello(context.Background(), &proto_bak.HelloRequest{
		Name: "hobby",
		G:    proto_bak.Gender_FEMALE,
		Mp: map[string]string{
			"name":    "booby",
			"company": "慕课网",
		},
		AddTime: timestamppb.New(time.Now()),
	})
	fmt.Println(r.Message)
}
