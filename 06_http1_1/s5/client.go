package main

import (
	"log"
	"net/rpc/jsonrpc"
)

// 引数
type ClientArgs struct {
	A, B int
}

func main() {
	client, err := jsonrpc.Dial("tcp", "localhost:18888")
	if err != nil {
		panic(err)
	}
	var result int
	args := &ClientArgs{4, 5}
	err = client.Call("Calculator.Multiply", args, &result)
	if err != nil {
		panic(err)
	}
	log.Printf("4 x 5 = %d\n", result)
}
