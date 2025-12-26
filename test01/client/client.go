package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, _ := rpc.Dial("tcp", "localhost:1234")
	defer client.Close()
	var reply *string = new(string)
	err := client.Call("HelloService.Hello", "bobby", reply)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*reply)
}
