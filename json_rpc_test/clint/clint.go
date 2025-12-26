package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost:1234")
	var reply *string = new(string)
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	err := client.Call("HelloService.Hello", "bobby", reply)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*reply)
}
