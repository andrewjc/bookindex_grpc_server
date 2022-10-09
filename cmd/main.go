package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"index_service/internal/indexer"
	"index_service/internal/search"
	"log"
	"net"
	"os"
	"path"

	pb "index_service/internal/grpc/proto"
)

var CurWorkingDirectory, _ = os.Getwd()

var (
	app_config = map[string]interface{}{
		"app_name": "index_server",
		"port":     8080,
		"folder":   path.Join(CurWorkingDirectory, "document_root"),
	}
)

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", app_config["port"]))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	idx := indexer.IndexProcess(&app_config)

	go func() {
		// Start the indexer
		idx.Start()
	}()

	pb.RegisterBookIndexServer(s, search.NewIndexServer(&app_config))

	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
