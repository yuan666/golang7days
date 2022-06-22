package main

import (
	"context"
	grpc "google.golang.org/grpc"
	"log"
	"net"
	proto "proto"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(
	ctx context.Context, args *proto.String,
) (*proto.String, error) {
	reply := &proto.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func main() {
	grpcSvr := grpc.NewServer()
	proto.RegisterHelloServiceServer(grpcSvr, new(HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcSvr.Serve(lis)
}
