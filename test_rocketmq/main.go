package main

import (
	"context"
	"fmt"
	"os"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

// Package main implements a simple producer to send message.
func main() {
	p, _ := rocketmq.NewProducer(producer.WithNameServer([]string{"127.0.0.1:9876"}))
	err := p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		os.Exit(1)
	}

	res, err := p.SendSync(context.Background(), &primitive.Message{
		Topic: "imooc",
		Body:  []byte("this is imooc"),
	})
	if err != nil {
		fmt.Printf("发送失败%s\n", err)
	} else {
		fmt.Printf("发送成功%s\n", res.String())
	}

	p.Shutdown()

}
