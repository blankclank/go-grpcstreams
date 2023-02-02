package main

import (
	"fmt"
	hardwaremonitoring "grpcstreams/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Streaming HW monitoring")

	lis, err := net.Listen("tcp", ":7777")
	if err != nil {
		panic(err)
	}

	gRPCserver := grpc.NewServer()

	s := &Server{}

	hardwaremonitoring.RegisterHardwareMonitorServer(gRPCserver, s)

	log.Println(gRPCserver.Serve((lis)))
}
