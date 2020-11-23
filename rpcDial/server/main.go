package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//HelloService hello
type HelloService struct{}

//HelloWorld rpc func
func (h *HelloService) HelloWorld(in string, out *string) error {
	*out = fmt.Sprintf("Hello %s", in)
	return nil
}

func main() {
	rpc.RegisterName("HelloService", &HelloService{})

	listener, err := net.Listen("tcp", ":8099")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
	
		// go rpc.ServeConn(conn)
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

	
}
