package main

import "net/rpc"

const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

////////////////////////////////////////////////////////////////
type HelloserviceClient struct {
	*rpc.Client
}

var _ HelloServiceInterface = (*HelloserviceClient)(nil)

func DialHelloService(network, address string) (*HelloserviceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloserviceClient{Client: c}, nil
}

func (hc *HelloserviceClient) Hello(request string, reply *string) error {
	return hc.Client.Call(HelloServiceName+".Hello", request, reply)
}
/////////////////////////////////////////////////////////////////
type HelloServiceServer struct {
}
var _ HelloServiceInterface = (*HelloServiceServer)(nil)

func (hs *HelloServiceServer) Hello(request string, reply *string) error  {
	*reply = "hello:"+request
	return nil
}