package grpc

import (
	"context"

	summary "github.com/SaucySalamander/owl-db/pkg/proto"
	"google.golang.org/grpc"
)

type server struct {
	summary.UnimplementedGetSummaryServer
}

func (s *server) GetSummary(ctx context.Context, request *summary.SummaryRequest) (*summary.SummaryResponse, error) {
	return &summary.SummaryResponse{Message: "test"}, nil
}

func RegisterServer(s *grpc.Server) {
	summary.RegisterGetSummaryServer(s, &server{})
}
