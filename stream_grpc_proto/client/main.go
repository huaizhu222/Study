package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"Rpc.Study.go/stream_grpc_proto/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial("localhost:1234", grpc.WithInsecure())

	c := proto.NewGreeterClient(conn)
	res, _ := c.GetStream(context.Background(), &proto.StreamReqData{})
	for {
		a, err := res.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(a)
	}
	//
	putS, _ := c.PutStream(context.Background())
	for i := 0; i < 10; i++ {
		putS.Send(&proto.StreamReqData{
			Data: fmt.Sprintf("慕课网%d", i),
		})
		time.Sleep(time.Second)
	}

	//双向流模式
	allStr, _ := c.AllStream(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			allStr.Send(&proto.StreamReqData{
				Data: "慕课网",
			})
			time.Sleep(time.Second)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			if a, err := allStr.Recv(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("接收到服务端消息:" + a.Data)
			}
		}
	}()
	wg.Wait()

}
