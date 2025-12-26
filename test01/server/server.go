package main

import (
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello," + request
	return nil
}
func main() {
	listener, _ := net.Listen("tcp", ":1234")
	defer listener.Close()
	rpc.RegisterName("HelloService", &HelloService{})

	for {
		conn, _ := listener.Accept()
		go func() {
			defer conn.Close()
			rpc.ServeConn(conn)
		}()
	}
}
