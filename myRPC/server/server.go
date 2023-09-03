package main

import (
	"log"
	"net"
	"net/rpc"
)

type MathService struct {
}

type Args struct {
	A int
	B int
}

type MyInterface interface {
	Multiply(*Args, *int) error
}

func RegisterServer(i *MathService) {
	err := rpc.Register(i)
	if err != nil {
		log.Fatalln("rpc.Register error:", err)
	}
}

func (m *MathService) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func Serve() {
	math := new(MathService)
	RegisterServer(math)

	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Listen error:", err)
	}
	log.Println("RPC server is listen on port 8080...")

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatalln("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}

func main() {
	Serve()
}
