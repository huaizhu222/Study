package main

import (
	"context"
	"fmt"
	"time"

	"Rpc.Study.go/grpc_test01/proto"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func main() {
	interceptor := func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Println(time.Since(start))
		return err
	}
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithUnaryInterceptor(interceptor))
	grpc_rtry := []retry.CallOption{
		retry.WithMax(3),
		retry.WithPerRetryTimeout(1 * time.Second),
		retry.WithCodes(codes.Unknown, codes.DeadlineExceeded, codes.Unavailable),
	}
	
	opts = append(opts, grpc.WithUnaryInterceptor(retry.UnaryClientInterceptor(grpc_rtry...)))
	conn, _ := grpc.Dial("127.0.0.1:1234", opts...)
	defer conn.Close()
	c := proto.NewGreeterClient(conn)

	r, _ := c.SayHello(context.Background(), &proto.HelloRequest{
		Name: "hobby",
	})
	fmt.Println(r.Message)
}
