package main

import (
	"context"
	"fmt"
	hardwaremonitoring "grpcstreams/proto"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "localhost:7777", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := hardwaremonitoring.NewHardwareMonitorClient(conn)

	emptyreq := &hardwaremonitoring.EmptyRequest{}

	stream, err := client.Monitor(ctx, emptyreq)
	if err != nil {
		panic(err)
	}

	stop := time.NewTicker(7 * time.Second)

	for {
		select {
		case <-stop.C:
			err := stream.CloseSend()
			if err != nil {
				log.Fatal("Failed to close stream: ", err.Error())
			}
			return
		default:
			res, err := stream.Recv()
			if err != nil {
				panic(err)
			}

			fmt.Println("New Hardware stat received")
			fmt.Println("CPU Usage: ", res.Cpu)
			fmt.Println("Memory Used: ", res.MemoryUsed)
			fmt.Println("Memory Free: ", res.MemoryFree)
		}
	}
}
