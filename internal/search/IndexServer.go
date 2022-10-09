package search

import pb "index_service/internal/grpc/proto"

type IndexServer struct {
	*pb.UnimplementedBookIndexServer
}

func (s *IndexServer) Search(req *pb.SearchRequest, stream pb.BookIndex_SearchServer) error {
	stream.Send(&pb.SearchResult{Message: "test"})

	return nil
}
