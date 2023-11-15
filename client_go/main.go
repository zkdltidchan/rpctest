package main

import (
	"fmt"
	"log"
	"net/rpc"
)

var ServerName = "TestServer"
var ServerFunction = "SayHello"

// Same as server
type Request struct {
	Name string
}

type Response struct {
	Message string
}

func main() {
	port := ":1234"
	// connect to rpc server
	client, err := rpc.Dial("tcp", "localhost"+port)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	req := &Request{
		Name: "Client",
	}

	res := &Response{}

	err = client.Call(
		ServerName+"."+ServerFunction,
		req,
		res,
	)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Response: %s\n", res.Message)
}
