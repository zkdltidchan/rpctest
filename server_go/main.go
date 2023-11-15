package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

// Request struct
type Request struct {
	Name string
}

type Response struct {
	Message string
}

// Server struct
type TestServer struct {
}

// SayHello func
func (t *TestServer) SayHello(request *Request, response *Response) error {
	response.Message = "Hello " + request.Name
	return nil
}

func main() {
	port := ":1234"

	testServer := new(TestServer)

	// register rpc server
	rpc.Register(testServer)

	// listen tcp port
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	fmt.Println("listen on" + port + "...")
	// block and wait for rpc request
	rpc.Accept(listen)
}
