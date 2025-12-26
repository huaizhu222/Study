package main

import (
	"fmt"
	"log"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/flow"
)

func main() {
	err := sentinel.InitDefault()
	if err != nil {
		log.Fatalf("初始化sentinel 异常:%v", err)
	}

	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "some-test",
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Throttling, // 流控效果为直接拒绝
			Threshold:              0.1,             // 请求的间隔控制在 1000/10=100 ms
			MaxQueueingTimeMs:      100,             // 最长排队等待时间
		},
	})

	if err != nil {
		log.Fatalf("初始化rule 异常:%v", err)
	}
	for i := 1; i <= 12; i++ {
		e, b := sentinel.Entry("some-test", sentinel.WithTrafficType(base.Inbound))
		if b != nil {
			fmt.Println("限流了")
		} else {
			fmt.Println("检查通过")
			e.Exit()
		}
	}

}
