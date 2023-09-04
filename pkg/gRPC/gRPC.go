package grpc

import (
	"context"

	internal_account "internal/accounts"

	"github.com/SaucySalamander/owl-db/pkg/proto"
	"google.golang.org/grpc"
)

type account_server struct {
	proto.UnimplementedAccountsServer
}

type summary_server struct {
	proto.UnimplementedGetSummaryServer
}

func (s *summary_server) GetSummary(ctx context.Context, request *proto.SummaryRequest) (*proto.SummaryResponse, error) {
	return &proto.SummaryResponse{Message: "test"}, nil
}

func (s *account_server) CreateAccount(ctx context.Context, request *proto.CreateAccountRequest) (*proto.CreateAccountResponse, error) {
	result := internal_account.CreateAccount(request)
	var response proto.CreateAccountResponse
	r, _ := result.RowsAffected()
	if r < 1 {
		response = proto.CreateAccountResponse{Response: false}
	} else {
		response = proto.CreateAccountResponse{Response: true}
	}

	return &response, nil
}

func RegisterServer(s *grpc.Server) {
	proto.RegisterGetSummaryServer(s, &summary_server{})
	proto.RegisterAccountsServer(s, &account_server{})
}
