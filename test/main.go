package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type Test struct {
	Id int `json:"id" query:"id"`
}

var mutex sync.Mutex
var mp sync.Map
var wg sync.WaitGroup

func main() {
	go func() {
		ch := make(chan int)
		ch <- 1
	}()
	time.Sleep(time.Second)
	a := gin.Default()
	a.GET("helloworld", func(ctx *gin.Context) {
		id := Test{}
		ctx.ShouldBind(&id)
		fmt.Println(id)
		ctx.JSON(http.StatusOK, id)
	})
	a.Run("127.0.0.1:8080")
	http.ListenAndServe(":8091", nil)

	// runtime.Gosched()
	// mp := make(map[string]string)
	// mp["2"] =  "1"
	// mp.Store("a", 1)
	// // fmt.Println(mp.Load("a"))
	// if t, ok := mp.Swap("a", 2); ok {
	// 	fmt.Println(t)
	// }
	// fmt.Println(mp.Load("a"))
	// ch := make(chan int, 1)
	// ch <- 1
	// // ctx ,cancel := context.WithCancel(context.Background())
	// ch2 := make(chan int, 3)
	// select {
	// case <-ch:
	// 	fmt.Print("222")
	// case <-ch2:
	// 	fmt.Print(2)
	// case <-time.After(time.Second * 1):
	// 	fmt.Println("timeout 1")
	// }
	// mp := make(chan int, 2)
	// mp <- 1

	// select {
	// case <-ch:
	// 	fmt.Println()
	// case <-time.After(time.Millisecond * time.Duration(10)):
	// }

}
