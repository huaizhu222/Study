package main

import (
	"context"
	"fmt"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

// Package main implements a simple producer to send message.
func main() {
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"127.0.0.1:9876"}), consumer.WithGroupName("mxshop"))

	err := c.Subscribe("Transaction", consumer.MessageSelector{}, func(ctx context.Context, megs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range megs {
			fmt.Printf("获取到%v \n", megs[i])
		}
		return consumer.ConsumeSuccess, nil
	})

	if err != nil {
		fmt.Println("订阅失败")
	}
	c.Start()

	time.Sleep(time.Hour)

	c.Shutdown()

}
