package main

import (
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	// jaegerlog "github.com/uber/jaeger-client-go/log"
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
		ServiceName: "mxshop-1",
	}

	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(err)
	}
	defer closer.Close()

	span := tracer.StartSpan("funcA")
	time.Sleep(time.Microsecond * 500)
	span.Finish()

	cfg2 := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "120.55.0.93:6831",
		},
		ServiceName: "mxshop-2",
	}

	tracer2, closer2, err2 := cfg2.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err2 != nil {
		panic(err)
	}
	defer closer2.Close()
	span1 := tracer2.StartSpan("funcB", opentracing.ChildOf(span.Context()))
	time.Sleep(time.Microsecond * 1000)
	defer span1.Finish()
}
