package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8099")
	if err != nil {
		log.Fatal(err)
	}
	var out string

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))


	in:=os.Args[1]
	err = client.Call("HelloService.HelloWorld", in, &out)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(out)
}
