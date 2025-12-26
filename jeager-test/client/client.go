package main

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"

	"Rpc.Study.go/jeager-test/otgrpc"
	"Rpc.Study.go/jeager-test/proto"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"
)

func main() {
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "120.55.0.93:6831",
		},
		ServiceName: "mxshop",
		
	}

	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(err)
	}
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	conn, _ := grpc.Dial("127.0.0.1:1234", grpc.WithInsecure(), grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())))
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	r, _ := c.SayHello(context.Background(), &proto.HelloRequest{
		Name: "hobby",
	})
	fmt.Println(r.Message)
}
