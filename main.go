package main

import (
	"fmt"
	hardwaremonitoring "grpcstreams/proto"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
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

	go func() {
		log.Fatal(gRPCserver.Serve(lis))
	}()

	grpcWebServer := grpcweb.WrapServer(gRPCserver)

	multiplex := grpcMultiplexer{
		grpcWebServer,
	}

	r := http.NewServeMux()
	webapp := http.FileServer(http.Dir("webapp/hwmonitor/build"))
	r.Handle("/", multiplex.Handler(webapp))

	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

type grpcMultiplexer struct {
	*grpcweb.WrappedGrpcServer
}

func (m *grpcMultiplexer) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if m.IsGrpcWebRequest(r) {
			m.ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
