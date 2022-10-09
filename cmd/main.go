package main

import (
    "flag"
    "fmt"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "index_service/internal/search"
    "log"
    "net"

    pb "index_service/internal/grpc/proto"
)

var (
    port = flag.Int("port", 8080, "The server port")
)

func main() {

    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()

    pb.RegisterBookIndexServer(s, &search.IndexServer{})

    reflection.Register(s)

    log.Printf("server listening at %v", lis.Addr())
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
