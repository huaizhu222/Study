package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"user_srv/handler"
	"user_srv/initialize"
	"user_srv/proto"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func main() {
	app := fx.New(
		// 模块注册
		fx.Provide(
			initialize.NewLogger,   // 日志组件
			initialize.NewConfig,    // 配置组件
			initialize.NewDB,        // 数据库组件
			NewJaegerTracer,         // 链路追踪组件
			NewGRPCServer,           // gRPC服务器组件
			NewHTTPServer,           // HTTP服务器组件（用于pprof）
		),
		fx.Invoke(
			func RegisterGRPCServices(server *grpc.Server) {
	proto.RegisterUserServer(server,  &handler.User_Server{})
	proto.RegisterSystemServer(server, &handler.System_Server{})
	zap.L().Info("GRPC services registered successfully")
}
,    // 注册gRPC服务
			StartServers,            // 启动服务器
		),
		fx.WithLogger(func(logger *zap.Logger) fx.Option {
			return fx.Logger(zap.NewStdLog(logger)) // 集成zap日志
		}),
	)

	// 启动应用
	if err := app.Start(context.Background()); err != nil {
		log.Fatalf("Failed to start application: %v", err)
	}

	// 等待中断信号
	<-app.Done()

	// 停止应用
	if err := app.Stop(context.Background()); err != nil {
		log.Fatalf("Failed to stop application gracefully: %v", err)
	}
}

// NewJaegerTracer 创建Jaeger追踪器（带生命周期管理）
func NewJaegerTracer(lc fx.Lifecycle) (opentracing.Tracer, error) {
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "120.55.0.93:6831",
		},
		ServiceName: "user-srv",
	}

	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		return nil, fmt.Errorf("failed to create Jaeger tracer: %w", err)
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			closer.Close()
			zap.L().Info("Jaeger tracer closed")
			return nil
		},
	})

	return tracer, nil
}

// NewGRPCServer 创建gRPC服务器（带生命周期管理）
func NewGRPCServer(tracer opentracing.Tracer, lc fx.Lifecycle) *grpc.Server {
	server := grpc.NewServer(
		grpc.UnaryInterceptor(otgrpc.OpenTracingServerInterceptor(tracer)),
	)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			lis, err := net.Listen("tcp", "127.0.0.1:50051")
			if err != nil {
				return fmt.Errorf("failed to listen: %w", err)
			}

			go func() {
				zap.L().Info("Starting gRPC server", zap.String("addr", "127.0.0.1:50051"))
				if err := server.Serve(lis); err != nil {
					zap.L().Error("gRPC server failed", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			zap.L().Info("Stopping gRPC server gracefully")
			server.GracefulStop()
			return nil
		},
	})

	return server
}

// NewHTTPServer 创建HTTP服务器（用于pprof）
func NewHTTPServer(lc fx.Lifecycle) *http.Server {
	server := &http.Server{
		Addr: ":6060", // 使用不同端口避免冲突
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				zap.L().Info("Starting HTTP server", zap.String("addr", ":6060"))
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					zap.L().Error("HTTP server failed", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			zap.L().Info("Stopping HTTP server")
			return server.Shutdown(ctx)
		},
	})

	return server
}

// RegisterGRPCServices 注册gRPC服务实现

// StartServers 启动服务（依赖注入触发）
func StartServers(
	_ *grpc.Server,  // 触发gRPC服务创建
	_ *http.Server,  // 触发HTTP服务创建
) {
	// 此处参数仅用于触发依赖关系，实际启动逻辑在生命周期钩子中
}