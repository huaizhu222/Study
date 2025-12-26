package main

import (
	"fmt"
	"log"
	"time"

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
			TokenCalculateStrategy: flow.WarmUp,
			ControlBehavior:        flow.Reject, // 流控效果为直接拒绝
			Threshold:              1000,        // 请求的间隔控制在 1000/10=100 ms
			// StatIntervalInMs: 1000,
			// WarmUpColdFactor: 1,
			WarmUpPeriodSec: 30,
		},
	})
	golbaltotal := 0
	passtotal := 0
	rebacktotal := 0
	if err != nil {
		log.Fatalf("初始化rule 异常:%v", err)
	}
	for i := 0; i < 3; i++ {
		go func() {
			for {
				e, b := sentinel.Entry("some-test", sentinel.WithTrafficType(base.Inbound))
				golbaltotal++
				if b != nil {
					// fmt.Println("限流了")
					rebacktotal++
				} else {
					// fmt.Println("检查通过")
					passtotal++
					e.Exit()
				}
			}
		}()
	}
	time.Sleep(30 * time.Second)
	fmt.Println(golbaltotal)
	fmt.Println(passtotal)
	fmt.Println(rebacktotal)

}
