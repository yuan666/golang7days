package main

import (
	"log"
	"net"
	"net/rpc"
)

func main() {
	RegisterHelloService(new(HelloServiceServer))

	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Listen tcp error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}
/****
	RegisterHelloService  函数来注册函数，避免了命名空间工作，保证了传入的参数，符合RPC的定义
	使用for来支持多个TCP链接，然后为每个链接提供RPC服务
缺点；
	net/rpc 使用gob编码，不不支持跨语言，若想使用跨语言，可考虑采用 net/rpc/jsonrpc

编译命令：go build  server.go svrclimethod.go
**/