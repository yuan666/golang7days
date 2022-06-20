package main

import (
	"fmt"
	"log"
)

func main() {
	/*
		client, err := rpc.Dial("tcp", "localhost:1234")
		if err != nil {
			log.Fatal("dialing:", err)

			var reply string

			err = client.Call(HelloServiceName+".Hello", "request hello",&reply)
			if err != nil {
				log.Fatal(err)
			}
		}
	*/
	//优化后的代码
	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello("request hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
/***
	DialHelloService 创建rpc客户端，入参为 使用的协议类型，监听的地址，出参包含 client
	然后调用client的Hello方法，这里讲tpc的函数都封装了

编译命令：go build client.go svrclimethod.go
**/