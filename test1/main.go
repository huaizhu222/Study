package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	Smap := sync.Map{}
	// ctx,cancel := context.WithCancel(context.Background())
	Smap.Store("key1", "val1")
	v, ok := Smap.Load("key1")
	if !ok {
		return
	}
	fmt.Println(v)

	var pInt atomic.Pointer[int]
	p := 42
	pInt.Store(&p)
	p = 23

}
