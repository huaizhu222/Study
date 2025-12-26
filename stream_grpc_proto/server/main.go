package main

import (
	"fmt"
	"net"
	"sync"
	"time"

	"Rpc.Study.go/stream_grpc_proto/proto"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedGreeterServer // 必须嵌入这个未实现的结构体以符合接口要求
}

func (s *server) PutStream(cliStr grpc.ClientStreamingServer[proto.StreamReqData, proto.StreamResData]) error {
	for {
		if a, err := cliStr.Recv(); err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(a.Data)
		}
	}
	return nil
}
func (s *server) AllStream(allStr grpc.BidiStreamingServer[proto.StreamReqData, proto.StreamResData]) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			allStr.Send(&proto.StreamResData{
				Data: fmt.Sprintf("%v", time.Now().Unix()),
			})
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			if a, err := allStr.Recv(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("接收到客户端消息:" + a.Data)
			}
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	return nil

}
func (s *server) GetStream(req *proto.StreamReqData, res grpc.ServerStreamingServer[proto.StreamResData]) error {
	i := 1
	for {
		i++
		res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}

func main() {
	lis, _ := net.Listen("tcp", "localhost:1234")
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}
