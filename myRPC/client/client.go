package main

import (
	"log"
	"net/rpc"
)

type Args struct {
	A int
	B int
}

func Multiply() {
	client, err := rpc.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalln("Dial error:", err)
	}
	args := &Args{A: 5, B: 20}
	var result int
	err = client.Call("MathService.Multiply", args, &result)

	if err != nil {
		log.Fatalln("Call error:", err)
	}
	log.Println("Result:", result)
}

func main() {
	Multiply()
}
