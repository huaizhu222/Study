package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

type OrderListener struct{}

func (*OrderListener) ExecuteLocalTransaction(*primitive.Message) primitive.LocalTransactionState {
	fmt.Println("开始执行本地逻辑")
	time.Sleep(time.Second * 3)
	fmt.Println("执行本地逻辑失败")
	return primitive.UnknowState

}
func (*OrderListener) CheckLocalTransaction(*primitive.MessageExt) primitive.LocalTransactionState {
	fmt.Println("消息回查")
	time.Sleep(time.Second * 15)
	return primitive.CommitMessageState
}

func main() {
	
	p, _ := rocketmq.NewTransactionProducer(&OrderListener{}, producer.WithNameServer([]string{"127.0.0.1:9876"}))
	err := p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		os.Exit(1)
	}

	res, err := p.SendMessageInTransaction(context.Background(), primitive.NewMessage("Transaction", []byte("This is Transaction 100")))
	if err != nil {
		fmt.Printf("发送失败%s\n", err)
	} else {
		fmt.Printf("发送成功%s\n", res.String())
	}

	time.Sleep(time.Hour)
	if err := p.Shutdown(); err != nil {
		panic("生成produer失败")
	}

}
