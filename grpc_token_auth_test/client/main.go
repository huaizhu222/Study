package main

import (
	"context"
	"fmt"

	"Rpc.Study.go/grpc_test01/proto"
	"google.golang.org/grpc"
)

type customCredential struct {
	token string
}

func (s *customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":    "bobby",
		"password": "imooc",
	}, nil
}
func (s *customCredential) RequireTransportSecurity() bool {
	return false
}

func main() {
	// interceptor := func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// 	start := time.Now()
	// 	md := metadata.New(map[string]string{
	// 		"appid":    "bobby",
	// 		"password": "imooc",
	// 	})
	// 	ctx = metadata.NewOutgoingContext(ctx, md)
	// 	err := invoker(ctx, method, req, reply, cc, opts...)
	// 	fmt.Println(time.Since(start))
	// 	return err
	// }
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithPerRPCCredentials(&customCredential{}))

	conn, _ := grpc.Dial("127.0.0.1:1234", opts...)
	defer conn.Close()
	c := proto.NewGreeterClient(conn)

	r, _ := c.SayHello(context.Background(), &proto.HelloRequest{
		Name: "hobby",
	})

	fmt.Println(r.Message)
}
