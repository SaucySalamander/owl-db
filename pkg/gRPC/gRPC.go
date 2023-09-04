package grpc

import (
	"context"

	internal_account "internal/accounts"

	"github.com/SaucySalamander/owl-db/pkg/proto/account"
	"github.com/SaucySalamander/owl-db/pkg/proto/summary"
	"google.golang.org/grpc"
)

type account_server struct {
	account.UnimplementedAccountsServer
}

type summary_server struct {
	summary.UnimplementedGetSummaryServer
}

func (s *summary_server) GetSummary(ctx context.Context, request *summary.SummaryRequest) (*summary.SummaryResponse, error) {
	return &summary.SummaryResponse{Message: "test"}, nil
}

func (s *account_server) CreateAccount(ctx context.Context, request *account.CreateAccountRequest) (*account.CreateAccountResponse, error) {
	result := internal_account.CreateAccount(request)
	var response account.CreateAccountResponse
	r, _ := result.RowsAffected()
	if r < 1 {
		response = account.CreateAccountResponse{Response: false}
	} else {
		response = account.CreateAccountResponse{Response: true}
	}

	return &response, nil
}

func RegisterServer(s *grpc.Server) {
	summary.RegisterGetSummaryServer(s, &summary_server{})
	account.RegisterAccountsServer(s, &account_server{})
}
